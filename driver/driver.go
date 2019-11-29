package driver

import "github.com/sawadashota/gorew/driver/configuration"

type Driver interface {
	Configuration() configuration.Provider
}

type DefaultDriver struct {
	c configuration.Provider
}

// NewDefaultDriver .
func NewDefaultDriver() Driver {
	c := configuration.NewEnvConfigProvider()

	return &DefaultDriver{
		c: c,
	}
}

// Configuration .
func (d *DefaultDriver) Configuration() configuration.Provider {
	return d.c
}
