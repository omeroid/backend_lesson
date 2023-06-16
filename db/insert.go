package db

import (
	"gorm.io/gorm"
)

// テーブルにレコードを挿入する
func InsertSampleRecord(db *gorm.DB) {
	//各テーブルに挿入するサンプルレコードの構造体を生成
	user := User{
		Name:         "wada hiroka",
		PasswordHash: "au923o",
	}
	message := Message{
		RoomID: 1,
		UserID: 1,
		Text:   "adadada",
	}

	description := "aaaa"
	room := Room{
		Description: &description,
		Name:        "chat room",
	}

	session := Session{
		UserID: 1,
		Token:  "aaaihjja",
	}

	//生成した構造体をDBにinsertする
	if result := db.Create(&user); result.Error != nil {
		panic("Failed to insert user")
	}

	if result := db.Create(&room); result.Error != nil {
		panic("Failed to insert room")
	}

	if result := db.Create(&message); result.Error != nil {
		panic("Failed to insert message")
	}

	if result := db.Create(&session); result.Error != nil {
		panic("Failed to insert session")
	}

}
