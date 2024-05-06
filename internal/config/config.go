package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	App struct {
		Repo     string `yaml:"repo"`
		RestPort string `yaml:"port"`
	}
	DB DBConfig `yaml:"db"`
}

type DBConfig struct {
	Host     string `yaml:"host" env:"DB_HOST"`
	Port     int    `yaml:"port" env:"DB_PORT"`
	Username string `yaml:"username" env:"DB_USERNAME"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	Database string `yaml:"database" env:"DB_DATABASE"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {

	instance = &Config{}
	if err := cleanenv.ReadConfig("./config/config.yaml", instance); err != nil {
		log.Fatal(err)
	}
	fmt.Println(instance)
	return instance
}
