package app

import (
	"github.com/ezraisw/tanshogyo/pkg/common/adapter/grpc"
	"github.com/ezraisw/tanshogyo/pkg/common/adapter/web"
)

type Runners struct {
	GRPCRunner *grpc.GRPCRunner
	WebRunner  *web.WebRunner
}
