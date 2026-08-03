package config

type PostgresConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     uint   `mapstructure:"port" validate:"required"`
	Username string `mapstructure:"username" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	DBName   string `mapstructure:"db_name" validate:"required"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     uint   `mapstructure:"port" validate:"required"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password" validate:"required"`
	DBName   string `mapstructure:"db_name" validate:"required"`
}

type ClientConfig struct {
	Postgres *PostgresConfig `mapstructure:"postgres" validate:"required"`
	Redis    *RedisConfig    `mapstructure:"redis" validate:"required"`
}
