package main

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	ID          uint
	Title       string
	Description *string
	IsCompleted bool `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func connect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}
