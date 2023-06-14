package handler

import (
// "github.com/volatiletech/null/v8"
)

type User OutputCreateUser
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
type InputCreateUser struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type OutputCreateUser struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

type InputCheckUser struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type OutputCheckUser struct {
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

type OutputGetRoomDetailList struct {
	Rooms []RoomDetail
}

type InputCreateRoom struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type OutputCreateRoom struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"createdAt"`
}

type OutputGetRoomDetail struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"createdAt"`
}

type InputCreateMessage struct {
	UserID string `json:"userId"`
	Text   string `json:"text"`
}

type OutputCreateMessage Message

type OutputGetMessageDetailList struct {
	Messages []Message `json:"messages"`
}

type OutputDeleteMessage Message
