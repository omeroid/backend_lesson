package handler

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type Message struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
}

type Room struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreateUserInput struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type CreateUserOutput struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type CheckUserInput struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type CheckUserOutput struct {
	UserID   int    `json:"userId"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

type GetRoomDetailListOutput struct {
	Rooms []Room
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

type GetRoomDetailOutput struct {
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
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetMessageDetailListOutput struct {
	Messages []Message `json:"messages"`
}

type DeleteMessageOutput struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
}
