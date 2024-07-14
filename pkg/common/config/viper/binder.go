package viper

import (
	"strings"

	"github.com/ezraisw/tanshogyo/pkg/common/config"
	"github.com/spf13/viper"
)

type ViperBinder struct {
	v *viper.Viper
}

func NewViperBinder(properties config.BinderProperties) *ViperBinder {
	v := viper.New()

	for _, path := range properties.Paths {
		v.AddConfigPath(path)
	}
	v.SetConfigName(properties.FileName)
	v.SetEnvPrefix(properties.EnvPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.SetConfigType("yml")

	return &ViperBinder{v: v}
}

func (b ViperBinder) BindTo(i any) error {
	if err := b.v.ReadInConfig(); err != nil {
		return err
	}
	return b.v.Unmarshal(i)
}
