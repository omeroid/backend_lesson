package db

import (
	"errors"
	"gorm.io/gorm"
)

// データベースにテーブルを作成する
func Migrate(db *gorm.DB) error {
	// 前回起動時に作成したテーブルの全削除
	db.Migrator().DropTable(&Message{})
	db.Migrator().DropTable(&Session{})
	db.Migrator().DropTable(&Room{})
	db.Migrator().DropTable(&User{})

	// テーブルを作成する
	if err := db.AutoMigrate(&Room{}, &User{}, &Session{}, &Message{}); err != nil {
		return errors.New("failed to migrate")
	}

	// テーブルに外部キー制約を設定する
	if err := db.Migrator().CreateConstraint(&Message{}, "UserID"); err != nil {
		return errors.New("failed to create foreign key constraint for Message")
	}

	if err := db.Migrator().CreateConstraint(&Message{}, "RoomID"); err != nil {
		return errors.New("failed to create foreign key constraint for Message")
	}

	if err := db.Migrator().CreateConstraint(&Room{}, "RoomID"); err != nil {
		return errors.New("failed to create foreign key constraint for Room")
	}

	if err := db.Migrator().CreateConstraint(&Session{}, "UserID"); err != nil {
		return errors.New("failed to create foreign key constraint for Session")
	}

	return nil
}
