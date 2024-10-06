package storage

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type StorageConfig struct {
	Driver     string `yaml:"driver" env-default:"mysql"`
	Connection struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
		Params   struct {
			Charset   string `yaml:"charset"`
			ParseTime string `yaml:"parseTime"`
			Loc       string `yaml:"loc"`
		} `yaml:"params"`
	} `yaml:"connection"`
}

func (sc *StorageConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		sc.Connection.Username,
		sc.Connection.Password,
		sc.Connection.Host,
		sc.Connection.Port,
		sc.Connection.Database,
		sc.Connection.Params.Charset,
		sc.Connection.Params.ParseTime,
		sc.Connection.Params.Loc,
	)
}
