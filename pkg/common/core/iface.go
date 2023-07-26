package core

import (
	"context"

	"github.com/pwnedgod/tanshogyo/pkg/common/util/reflhelper"
)

type Runner interface {
	Run(context.Context) error
	Stop(context.Context) error
}

type HandlerRegistrar interface {
	AddHandlerRegistries(*reflhelper.StructCatalog) error
}

type MiddlewareRegistrar interface {
	AddMiddlewareRegistries(*reflhelper.StructCatalog) error
}
