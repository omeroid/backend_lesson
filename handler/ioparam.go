package handler

import ()

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
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}
type OutputGetRoomsDetail struct {
	Rooms []RoomDetail
}
