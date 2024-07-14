package config

import "github.com/ezraisw/tanshogyo/pkg/common/config"

func ProvideConfig(binder config.Binder) (*Config, error) {
	config := &Config{}
	if err := binder.BindTo(config); err != nil {
		return nil, err
	}
	return config, nil
}
