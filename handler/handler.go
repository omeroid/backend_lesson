package handler

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/omeroid/kosen_backend_lesson/db"
)

type ResponseGetMessages struct {
	Response []db.Message
}

func SendMessage(c echo.Context) error {

	//msg := c.FormValue("message")
	//roomID := c.FormValue("roomID")

	return c.String(http.StatusOK, "message sended")
}

func GetMessages(c echo.Context) error {

	conn, err := db.Initdb()
	if err != nil {
		log.Fatalln("接続失敗！", err)
	}

	var messages []db.Message
	conn.Find(&messages) //全件取得

	res := &ResponseGetMessages{}
	res.Response = messages

	var tmp []byte

	tmp, err = json.Marshal(res)

	return c.String(http.StatusOK, string(tmp))
}

//columnはroomID,content,timestampで行きます
