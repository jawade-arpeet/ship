package config

import "ship/internal/constants"

type ServerConfig struct {
	RunEnv constants.Env `mapstructure:"run_env" validate:"required"`
	Port   uint          `mapstructure:"port" validate:"required"`
}
