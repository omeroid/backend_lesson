package db

import (
	"gorm.io/gorm"
)

func Insert(db *gorm.DB) {
	user := User{
		Name:         "wada hiroka",
		PasswordHash: "au923o",
	}
	message := Message{
		RoomID: 1,
		UserID: 1,
		Text:   "adadada",
	}
	room := Room{
		Description: "aaaa",
		Name:        "chat room",
	}

	session := Session{
		UserID: 1,
		Token:  "aaaihjja",
	}

	result := db.Create(&user)
	if result.Error != nil {
		panic("Failed to insert user")
	}

	result = db.Create(&room)
	if result.Error != nil {
		panic("Failed to insert room")
	}

	result = db.Create(&message)
	if result.Error != nil {
		panic("Failed to insert message")
	}

	result = db.Create(&session)
	if result.Error != nil {
		panic("Failed to insert session")
	}

}
