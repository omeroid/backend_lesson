package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// type Messages struct {
// MessageID string `gorm:"column:messageid"`
// RoomID    string `gorm:"column:roomid"`
// Content   string `gorm:"column:content"`
// TimeStamp string `gorm:"column:timestamp"`
// }
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("chatsystem.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	return db, nil
}
