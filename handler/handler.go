package handler

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"

	//"fmt"
	//"log"
	//"math/rand"
	"net/http"
	"strconv"
	//	"time"

	"github.com/labstack/echo/v4"
	"github.com/omeroid/kosen_backend_lesson/db"
)

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

// signup
func CreateUser(c echo.Context) error {
	p := new(InputCreateUser)
	if err := c.Bind(p); err != nil {
		return c.String(http.StatusBadRequest, ThrowError(err.Error()))
	}

	conn, err := db.InitDB()
	if err != nil {
		return c.String(http.StatusBadRequest, ThrowError(err.Error()))
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(p.Password), 10)

	user := db.User{
		Name:         p.Username,
		PasswordHash: string(hashedPassword), //Hash化する
	}

	conn.Create(&user)

	res := OutputCreateUser{
		ID:        strconv.Itoa(user.ID),
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
	}

	var output []byte
	output, err = json.Marshal(res)

	return c.String(http.StatusCreated, string(output))

}

func ThrowError(error string) string {
	res := ErrorResponse{
		Message: error,
	}

	var output []byte
	output, _ = json.Marshal(res) //ここどうしよう
	return string(output)
}
