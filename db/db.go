package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DBに接続する
func InitDB(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
