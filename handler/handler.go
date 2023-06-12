package handler

import (
	"encoding/json"
	"fmt"
	//"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	//"fmt"
	//"log"
	//"math/rand"
	"net/http"
	"strconv"

	//	"time"

	"github.com/google/uuid"
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

type InputCheckUser struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}
type OutputCheckUser struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
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

	result := conn.Create(&user)

	if result.Error != nil {
		return c.String(http.StatusBadRequest, ThrowError(result.Error.Error()))
	}

	output := OutputCreateUser{
		ID:        strconv.Itoa(user.ID),
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
	}

	var res []byte
	res, err = json.Marshal(output)
	if err != nil {
		return c.String(http.StatusBadRequest, ThrowError(err.Error()))
	}

	return c.String(http.StatusCreated, string(res))

}

func CheckUser(c echo.Context) error {
	p := new(InputCheckUser)
	if err := c.Bind(p); err != nil {
		return c.String(http.StatusBadRequest, ThrowError(err.Error()))
	}

	conn, err := db.InitDB()
	if err != nil {
		return c.String(http.StatusBadRequest, ThrowError(err.Error()))
	}

	user := db.User{}
	result := conn.Take(&user, "name=?", p.Username)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, ThrowError(result.Error.Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(p.Password))
	if err != nil {
		return c.String(http.StatusBadRequest, ThrowError(err.Error()))
	}

	token, _ := uuid.NewRandom()

	//Sessionへの追加
	session := db.Session{
		UserID:    user.ID,
		Token:     token.String(),
		ExpiredAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	result = conn.Create(&session)
	if result.Error != nil {
		c.String(http.StatusBadRequest, ThrowError(result.Error.Error()))
	}

	output := OutputCheckUser{
		UserID:   strconv.Itoa(user.ID),
		UserName: user.Name,
		Token:    token.String(),
	}

	var res []byte
	res, err = json.Marshal(output)
	if err != nil {
		c.String(http.StatusBadRequest, ThrowError(err.Error()))
	}

	return c.String(http.StatusOK, string(res))
}

func ThrowError(error string) string {
	res := ErrorResponse{
		Message: error,
	}

	var output []byte
	output, _ = json.Marshal(res) //ここどうしよう
	return string(output)
}
