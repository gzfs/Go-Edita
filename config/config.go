package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	PORT   string `mapstructure:"PORT"`
	DB_URL string `mapstructure:"DB_URL"`
}

func Load() (envConfig Env, err error) {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "production" {
		return Env{
			PORT:   os.Getenv("PORT"),
			DB_URL: os.Getenv("DB_URL"),
		}, nil
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&envConfig)

	if envConfig.PORT == "" {
		envConfig.PORT = ":3000"
	}

	if envConfig.DB_URL == "" {
		err = errors.New("DB_URL must be set")
		return
	}

	return
}
