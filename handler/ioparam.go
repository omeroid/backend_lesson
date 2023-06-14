package handler

import ()

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}
type Message struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
	User      User   `json:"user"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// LoginParameters
type CreateUserInput struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type CreateUserOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

type CheckUserInput struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type CheckUserOutput struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

type RoomDetail struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"createdAt"`
}

type GetRoomDetailListOutput struct {
	Rooms []RoomDetail
}

type CreateRoomInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type CreateRoomOutput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"createdAt"`
}

type GetRoomDetailOutput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"createdAt"`
}

type CreateMessageInput struct {
	UserID string `json:"userId"`
	Text   string `json:"text"`
}

type CreateMessageOutput Message

type GetMessageDetailListOutput struct {
	Messages []Message `json:"messages"`
}

type DeleteMessageOutput Message
