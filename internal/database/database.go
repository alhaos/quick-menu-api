package database

import (
	"context"
	"github.com/jackc/pgx"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
	Database string `yaml:"database"`
}

func New(conf Config) (*pgx.Conn, error) {

	connConfig := pgx.ConnConfig{
		Host:     conf.Host,
		Port:     uint16(conf.Port),
		Database: conf.Database,
		User:     conf.User,
		Password: conf.Pass,
	}

	connect, err := pgx.Connect(connConfig)
	if err != nil {
		return nil, err
	}

	err = connect.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return connect, nil
}
