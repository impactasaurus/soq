package cmd

import (
	"github.com/kelseyhightower/envconfig"
)

type Network struct {
	Port int `envconfig:"PORT" default:"80"`
}

type Config struct {
	Network Network
}

func MustGetConfiguration() *Config {
	c := &Config{}
	envconfig.MustProcess("", c)
	return c
}
