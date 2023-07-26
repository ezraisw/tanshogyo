package app

import (
	"github.com/pwnedgod/tanshogyo/pkg/common/adapter/grpc"
	"github.com/pwnedgod/tanshogyo/pkg/common/adapter/web"
)

type Runners struct {
	GRPCRunner *grpc.GRPCRunner
	WebRunner  *web.WebRunner
}
