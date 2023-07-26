package web

import "fmt"

type HTTPConfigGetter interface {
	GetHTTPConfig() *HTTPConfig
}

type HTTPConfig struct {
	Host string
	Port int
}

func (h HTTPConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}
