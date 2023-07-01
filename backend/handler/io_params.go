package handler

import "time"

type SignUpInput struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type SignUpOutput struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserOutput struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type MessageOutput struct {
	ID        int        `json:"id"`
	Text      string     `json:"text"`
	User      UserOutput `json:"user"`
	CreatedAt time.Time  `json:"createdAt"`
}

type RoomOutput struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type ListRoomOutput struct {
	Rooms []RoomOutput
}

type CreateRoomInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type CreateRoomOutput struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type GetRoomOutput struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type CreateMessageInput struct {
	UserID int    `json:"userId"`
	Text   string `json:"text"`
}

type CreateMessageOutput struct {
	ID        int        `json:"id"`
	Text      string     `json:"text"`
	User      UserOutput `json:"user"`
	CreatedAt time.Time  `json:"createdAt"`
}

type ListMessageOutput struct {
	Messages []MessageOutput `json:"messages"`
}

type DeleteMessageOutput struct {
	ID        int        `json:"id"`
	Text      string     `json:"text"`
	User      UserOutput `json:"user"`
	CreatedAt time.Time  `json:"createdAt"`
}
