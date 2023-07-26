package grpc

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/pwnedgod/tanshogyo/pkg/common/util/helper"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/reflhelper"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type GRPCRunnerOptions struct {
	Logger *zerolog.Logger
}

type GRPCRunner struct {
	o      GRPCRunnerOptions
	addr   string
	server *grpc.Server
}

func NewGRPCRunner(configGetter GRPCConfigGetter, options GRPCRunnerOptions) (*GRPCRunner, error) {
	config := configGetter.GetGRPCConfig()

	srvOptions, err := makeServerOptions(config)
	if err != nil {
		return nil, err
	}

	server := grpc.NewServer(srvOptions...)
	reflection.Register(server)

	return &GRPCRunner{
		o:      options,
		addr:   fmt.Sprintf("%s:%d", config.Host, config.Port),
		server: server,
	}, nil
}

func (r GRPCRunner) Run(context.Context) error {
	r.o.Logger.Info().
		Str("addr", r.addr).
		Msg("starting grpc server")

	lis, err := net.Listen("tcp", r.addr)
	if err != nil {
		return err
	}

	if err := r.server.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		return err
	}

	r.o.Logger.Info().
		Msg("exited grpc server")

	return nil
}

func (r GRPCRunner) Stop(ctx context.Context) error {
	r.o.Logger.Info().
		Msg("shutting down grpc server")

	w := helper.NewWaiter[struct{}]()
	defer func() { w.Close() }()

	w.Run(func() struct{} {
		r.server.GracefulStop()
		return struct{}{}
	})

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-w.Result():
	}

	return nil
}

func (r GRPCRunner) AddHandlerRegistries(catalog *reflhelper.StructCatalog) error {
	for _, c := range reflhelper.Collect[GRPCHandlerRegistry](catalog) {
		r.o.Logger.Debug().
			Str("name", c.Name).
			Msg("adding grpc handler registry")
		registry := c.Value
		if err := registry.RegisterServices(r.server); err != nil {
			return err
		}
	}

	serviceInfo := r.server.GetServiceInfo()
	r.o.Logger.Info().
		Int("count", len(serviceInfo)).
		Msg("services registered")

	return nil
}

func makeServerOptions(config *GRPCConfig) ([]grpc.ServerOption, error) {
	srvOptions := []grpc.ServerOption{}

	if config.CertificateFile != "" && config.PrivateKeyFile != "" {
		creds, err := credentials.NewServerTLSFromFile(config.CertificateFile, config.PrivateKeyFile)
		if err != nil {
			return nil, err
		}
		srvOptions = append(srvOptions, grpc.Creds(creds))
	}

	return srvOptions, nil
}
