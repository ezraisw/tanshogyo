package app

import (
	"github.com/pwnedgod/tanshogyo/pkg/common/adapter/web"
	"github.com/pwnedgod/tanshogyo/pkg/common/logger"
)

type MiddlewareRegistries struct {
	Essentials *web.EssentialsMiddlewareRegistry
	Logger     *logger.LoggerMiddlewareRegistry
}
