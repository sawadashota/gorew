package configuration

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	viperEnvPrefix = "gorew"

	viperLockfilePath   = "lockfile.path"
	defaultLockfilePath = "~/.gorew"
)

func init() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix(viperEnvPrefix)
	viper.AutomaticEnv()
}

type EnvConfigProvider struct{}

func NewEnvConfigProvider() *EnvConfigProvider {
	return new(EnvConfigProvider)
}

func (e *EnvConfigProvider) LockfilePath() string {
	path := viper.GetString(viperLockfilePath)
	if path == "" {
		path = defaultLockfilePath
	}
	return path
}
