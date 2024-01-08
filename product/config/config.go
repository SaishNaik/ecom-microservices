package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type App struct {
	Name    string `env-required="true" yaml:"name" env="APP_NAME"`
	Version string `env-required="true" yaml:"version" env="APP_VERSION"`
}

type HTTP struct {
	Host string `env-required="true" yaml:"host" env="HOST"`
	Port string `env-required="true" yaml:"port" env="PORT"`
}

type GRPC struct {
	Host string `env-required="true" yaml:"host" env="HOST"`
	Port string `env-required="true" yaml:"port" env="PORT"`
}

type PG struct {
	DNS string `env-required:"true" yaml:"dns" env:"PG_DNS"`
}

type Log struct {
	Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
}

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"log"`
		GRPC `yaml:"grpc"`
		PG   `yaml:"pg"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dir)

	err = cleanenv.ReadConfig(dir+"/config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error from file: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
