package db

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// テーブルにレコードを挿入する
func InsertSampleRecord(db *gorm.DB) error {
	//各テーブルに挿入するサンプルレコードの構造体を生成
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("backend"), 10)
	if err != nil {
		return errors.New("Failed to generate password hash")
	}
	user := User{
		Name:         "omeroid",
		PasswordHash: string(passwordHash),
	}
	message := Message{
		RoomID: 1,
		UserID: 1,
		Text:   "Welcome to the omeroid lecture!",
	}

	description := "どんな話題でもOK!　雑談ルーム"
	room := Room{
		Description: &description,
		Name:        "#general",
	}

	session := Session{
		UserID: 1,
		Token:  "57a336a5-d877-45aa-9fde-15c7c7309bec",
	}

	//生成した構造体をDBにinsertする
	if result := db.Create(&user); result.Error != nil {
		return errors.New("Failed to insert user")
	}

	if result := db.Create(&room); result.Error != nil {
		return errors.New("Failed to insert room")
	}

	if result := db.Create(&message); result.Error != nil {
		return errors.New("Failed to insert message")
	}

	if result := db.Create(&session); result.Error != nil {
		return errors.New("Failed to insert session")
	}

	return nil
}
