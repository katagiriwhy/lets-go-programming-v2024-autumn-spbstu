package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DataBase string `yaml:"dataBase"`
}

type ConfigDB struct {
	UserDB     string `yaml:"dbUser"`
	PasswordDB string `yaml:"dbPassword"`
	NameDB     string `yaml:"dbName"`
	HostDB     string `yaml:"dbHost"`
	PortDB     int    `yaml:"dbPort"`
}

func Load(name string) *ConfigDB {
	data, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	cfg := &ConfigDB{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return cfg
}

type Storage struct {
	Db *sql.DB
}

func NewStorage(cfg *ConfigDB) (*Storage, error) {
	conStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.UserDB,
		cfg.PasswordDB,
		cfg.HostDB,
		cfg.PortDB,
		cfg.NameDB,
	)
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Storage{Db: db}, nil
}
