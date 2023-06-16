package db

import (
	"time"
)

type User struct {
	ID           int        `gorm:"primaryKey;autoIncrement"`
	Name         string     `gorm:"column:name;NOT NULL;UNIQUE"`
	PasswordHash string     `gorm:"column:password_hash;NOT NULL"`
	Messages     []Message  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Sessions     []Session  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt    time.Time  `gorm:"column:created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at"`
}

type Message struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	RoomID    int       `gorm:"column:room_id;NOT NULL"`
	UserID    int       `gorm:"column:user_id;NOT NULL"`
	Text      string    `gorm:"column:text;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type Room struct {
	ID          int        `gorm:"primaryKey;autoIncrement"`
	Name        string     `gorm:"column:name;NOT NULL;"`
	Description *string    `gorm:"column:description"`
	Messages    []Message  `gorm:"foreignKey:RoomID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
}

type Session struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	UserID    int       `gorm:"column:user_id;NOT NULL"`
	Token     string    `gorm:"column:token;NOT NULL;UNIQUE"`
	ExpiredAt time.Time `gorm:"column:expired_at;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
