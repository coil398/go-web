package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Server *Server
	DB     *DBConfig
}

type Server struct {
	Port int
}

type DBConfig struct {
	User     string
	Password string
	Database string
	Host     string
	Port     int
	TimeZone string
}

func ReadDBConfig() (*DBConfig, error) {
	port, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		return nil, fmt.Errorf("Could not read the env var 'MYSQL_PORT': %w", err)
	}

	dbConfig := &DBConfig{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Database: os.Getenv("MYSQL_DATABASE"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     port,
		TimeZone: os.Getenv("TZ"),
	}

	return dbConfig, nil
}

func ReadConfig() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, fmt.Errorf("Could not read the env var 'SERVER_PORT': %w", err)
	}

	dbConfig, err := ReadDBConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		&Server{
			Port: port,
		},
		dbConfig,
	}, nil
}
