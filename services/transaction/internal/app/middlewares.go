package app

import (
	"github.com/ezraisw/tanshogyo/pkg/common/adapter/web"
	"github.com/ezraisw/tanshogyo/pkg/common/logger"
)

type MiddlewareRegistries struct {
	Essentials *web.EssentialsMiddlewareRegistry
	Logger     *logger.LoggerMiddlewareRegistry
}
