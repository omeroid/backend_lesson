package handler

import (
	"encoding/json"

	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/omeroid/kosen_backend_lesson/db"
)

var layout = "2006-01-02T15:04:05Z07:00"

type ResponseGetMessages struct {
	Response []db.Message
}

type InputSendMessage struct {
	RoomID  string `json:"roomID"`
	Message string `json:"message"`
}

func SendMessage(c echo.Context) error {

	p := new(InputSendMessage)
	if err := c.Bind(p); err != nil {
		log.Fatalln("パラメータが不正です", err)
		return err
	}

	conn, err := db.Initdb()
	if err != nil {
		log.Fatalln("接続失敗!", err)
	}

	now := time.Now()

	message := db.Message{
		RoomID:    p.RoomID,
		Content:   p.Message,
		TimeStamp: now.Format(layout),
	}

	conn.Create(&message)

	return c.String(http.StatusOK, "message sended")
}

func GetMessages(c echo.Context) error {

	roomID := c.Param("id")

	conn, err := db.Initdb()
	if err != nil {
		log.Fatalln("接続失敗！", err)
	}

	var messages []db.Message
	conn.Find(&messages, "roomid=?", roomID) //指定したroomIDのメッセージを全件取得

	res := &ResponseGetMessages{}
	res.Response = messages

	var output []byte

	output, err = json.Marshal(res)

	return c.String(http.StatusOK, string(output))
}

//columnはroomID,content,timestampで行きます
