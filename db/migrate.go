package db

import (
	"fmt"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	// Drop all tables
	db.Migrator().DropTable(&Message{})
	db.Migrator().DropTable(&Session{})
	db.Migrator().DropTable(&Room{})
	db.Migrator().DropTable(&User{})

	// Auto Migrate
	err := db.AutoMigrate(&Room{}, &User{}, &Session{}, &Message{})
	if err != nil {
		panic("failed to migrate")
	}

	// Set foreign key constraints
	if err := db.Migrator().CreateConstraint(&Message{}, "UserID"); err != nil {
		panic("failed to create foreign key constraint for Message")
	}

	if err := db.Migrator().CreateConstraint(&Message{}, "RoomID"); err != nil {
		panic("failed to create foreign key constraint for Message")
	}

	if err := db.Migrator().CreateConstraint(&Room{}, "RoomID"); err != nil {
		panic("failed to create foreign key constraint for Room")
	}

	if err := db.Migrator().CreateConstraint(&Session{}, "UserID"); err != nil {
		panic("failed to create foreign key constraint for Session")
	}

	fmt.Println("Migration successful")

}
