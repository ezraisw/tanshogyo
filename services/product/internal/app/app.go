package app

import (
	"context"

	"github.com/ezraisw/tanshogyo/pkg/common/core"
	"github.com/ezraisw/tanshogyo/pkg/common/entity"
	"github.com/ezraisw/tanshogyo/pkg/common/util/helper"
	"github.com/ezraisw/tanshogyo/pkg/common/util/reflhelper"
)

type ApplicationOptions struct {
	HandlerRegistries    HandlerRegistries
	MiddlewareRegistries MiddlewareRegistries
	Runners              Runners
	Migrator             entity.Migrator
}

type Application struct {
	o       ApplicationOptions
	runners []core.Runner
}

func NewApplication(options ApplicationOptions) *Application {
	collecteds := reflhelper.Collect[core.Runner](reflhelper.NewStructCatalog(options.Runners))
	runners := make([]core.Runner, 0, len(collecteds))
	for _, c := range collecteds {
		runners = append(runners, c.Value)
	}

	return &Application{
		o:       options,
		runners: runners,
	}
}

func (a Application) Run(ctx context.Context) error {
	if err := a.migrate(); err != nil {
		return err
	}

	if err := a.registerMiddlewares(); err != nil {
		return err
	}

	if err := a.registerHandlers(); err != nil {
		return err
	}

	if err := a.executeRunners(ctx); err != nil {
		return err
	}

	return nil
}

func (a Application) migrate() error {
	return a.o.Migrator.Migrate(models)
}

func (a Application) registerHandlers() error {
	handlerCatalog := reflhelper.NewStructCatalog(a.o.HandlerRegistries)
	for _, runner := range a.runners {
		if registrar, ok := runner.(core.HandlerRegistrar); ok {
			if err := registrar.AddHandlerRegistries(handlerCatalog); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a Application) registerMiddlewares() error {
	middlewareCatalog := reflhelper.NewStructCatalog(a.o.MiddlewareRegistries)
	for _, runner := range a.runners {
		if registrar, ok := runner.(core.MiddlewareRegistrar); ok {
			if err := registrar.AddMiddlewareRegistries(middlewareCatalog); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a Application) executeRunners(ctx context.Context) error {
	w := helper.NewWaiter[error]()
	defer func() { w.Close() }()

	for _, runner := range a.runners {
		r := runner
		w.Run(func() error {
			return r.Run(ctx)
		})
	}

	for range a.runners {
		if err := <-w.Result(); err != nil {
			return err
		}
	}

	return nil
}

func (a Application) Stop(ctx context.Context) error {
	for _, runner := range a.runners {
		if err := runner.Stop(ctx); err != nil {
			return err
		}
	}
	return nil
}
