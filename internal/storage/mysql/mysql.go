package mysql

import (
	"fmt"
	"url-shortener/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

type URL struct {
	ID    uint   `gorm:"primaryKey"`
	Alias string `gorm:"type:varchar(255);uniqueIndex"`
	URL   string `gorm:"type:text"`
}

func New(cfg *config.Config) (*Storage, error) {
	const op = "storage.mysql.New"

	db, err := gorm.Open(mysql.Open(cfg.Storage.DSN()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = db.AutoMigrate(&URL{})
	if err != nil {
		return nil, fmt.Errorf("%s: failed to migrate: %w", op, err)
	}

	return &Storage{db: db}, nil
}
