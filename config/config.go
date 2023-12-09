package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	PORT       string `mapstructure:"PORT"`
	DB_URL     string `mapstructure:"DB_URL"`
	API_KEY    string `mapstructure:"API_KEY"`
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
}

func Load() (envConfig Env, err error) {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "production" {
		return Env{
			PORT:       os.Getenv("PORT"),
			DB_URL:     os.Getenv("DB_URL"),
			API_KEY:    os.Getenv("API_KEY"),
			JWT_SECRET: os.Getenv("JWT_SECRET"),
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

	if envConfig.API_KEY == "" {
		err = errors.New("API_KEY must be set")
		return
	}

	if envConfig.JWT_SECRET == "" {
		err = errors.New("JWT_SECRET must be set")
		return
	}

	return
}
