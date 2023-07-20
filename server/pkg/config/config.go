package config

import (
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	Port               string `mapstructure:"PORT"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBName             string `mapstructure:"DB_NAME"`
	DBUser             string `mapstructure:"DB_USER"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBPassword         string `mapstructure:"DB_PASSWORD"`
	JwtKey             string `mapstructure:"JWT_KEY"`
	GoAuthClientId     string `mapstructure:"GOAUTH_CLIENT_ID"`
	GoAuthClientSecret string `mapstructure:"GOAUTH_CLIENT_SECRET"`
}

var envs = []string{
	"PORT", "DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "JWT_KEY",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
