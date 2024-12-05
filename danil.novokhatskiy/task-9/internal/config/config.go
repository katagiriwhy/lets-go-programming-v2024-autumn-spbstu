package config

import (
	"flag"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env         string        `yaml:"env"`
	StoragePath string        `yaml:"storage_path"`
	Addr        string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func ReadFlag() string {
	var path string
	flag.StringVar(&path, "config", "", "config file path")
	flag.Parse()
	return path
}

func LoadConfig(path string) *Config {
	if path == "" {
		log.Panic("Path can't be empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Panic("Config file doesn't exist")
	}
	var config Config
	file, err := os.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Panic(err)
	}
	return &config
}
