package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml: "env"`
	DB         `yaml: "db"`
	HTTPServer `yaml:"http_server"`
}

type DB struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
}

type HTTPServer struct {
	Address      string        `yaml: "Address"`
	Timeout      time.Duration `yaml: "timeout"`
	IddleTimeout time.Duration `yaml: "iddle_timeout"`
}

func MustLoad() *Config {
	configpath := "./config/local.yaml"
	if configpath == "" {
		log.Fatal("Config path is not set")
	}

	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist %s", err)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configpath, &cfg); err != nil {
		log.Fatalf("Cannot read config %s", err)
	}
	return &cfg
}
