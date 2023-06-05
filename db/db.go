package db

import (
	//"github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Message struct {
	RoomID    uint   `gorm:"column:roomid"`
	Content   string `gorm:"column:content"`
	TimeStamp string `gorm:"column:timestamp"`
}

func Initdb() (*gorm.DB, error) {
	dsn := "host=localhost user=root password=root dbname=root port=5430 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	return db, nil
}
