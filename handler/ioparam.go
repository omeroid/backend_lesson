package handler

import ()

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

type Message struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
	User      User   `json:"user"`
}

type Room struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"createdAt"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreateUserInput struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type CreateUserOutput struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
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
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"createdAt"`
}

type GetRoomDetailOutput struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"createdAt"`
}

type CreateMessageInput struct {
	UserID int    `json:"userId"`
	Text   string `json:"text"`
}

type CreateMessageOutput struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
	User      User   `json:"user"`
}

type GetMessageDetailListOutput struct {
	Messages []Message `json:"messages"`
}

type DeleteMessageOutput struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
	User      User   `json:"user"`
}
