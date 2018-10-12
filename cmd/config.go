package cmd

import (
	"github.com/kelseyhightower/envconfig"
)

type Network struct {
	Port int `envconfig:"PORT" default:"80"`
}

type Path struct {
	Questionnaires string `required:"true"`
}

type Config struct {
	Network Network
	Path    Path
}

func MustGetConfiguration() *Config {
	c := &Config{}
	envconfig.MustProcess("", c)
	return c
}
