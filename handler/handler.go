package handler

import (
	"encoding/json"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

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

// signup
func CreateUser(c echo.Context) error {
	p := new(InputCreateUser)
	if err := c.Bind(p); err != nil {
		return c.String(http.StatusBadRequest, ThrowError(err.Error()+" (入力値エラー)"))
	}

	conn, err := db.InitDB()
	if err != nil {
		return c.String(http.StatusBadRequest, ThrowError(err.Error()+" (DB接続エラー)"))
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(p.Password), 10)

	user := db.User{
		Name:         p.Username,
		PasswordHash: string(hashedPassword), //Hash化する
	}

	result := conn.Create(&user)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, ThrowError(result.Error.Error()+" (user作成エラー)"))
	}

	output := OutputCreateUser{
		ID:        strconv.Itoa(user.ID),
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
	}

	var res []byte
	res, err = json.Marshal(output)
	if err != nil {
		return c.String(http.StatusBadRequest, ThrowError(err.Error()+" (jsonのMarshalのエラー)"))
	}

	return c.String(http.StatusCreated, string(res))

}

func CheckUser(c echo.Context) error {
	p := new(InputCheckUser)
	if err := c.Bind(p); err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (入力値エラー)"))
	}

	conn, err := db.InitDB()
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (DB接続エラー)"))
	}

	user := db.User{}
	result := conn.Take(&user, "name=?", p.Username)
	if result.Error != nil {
		return c.String(http.StatusUnauthorized, ThrowError(result.Error.Error()+" (User検索エラー)"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(p.Password))
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (パスワードが違う)"))
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
		c.String(http.StatusUnauthorized, ThrowError(result.Error.Error()+" (session作成エラー)"))
	}

	output := OutputCheckUser{
		UserID:   strconv.Itoa(user.ID),
		UserName: user.Name,
		Token:    token.String(),
	}

	var res []byte
	res, err = json.Marshal(output)
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (jsonのMarshalのエラー)"))
	}

	return c.String(http.StatusOK, string(res))
}

func GetRoomDetailList(c echo.Context) error {

	conn, err := db.InitDB()
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (DBの接続エラー)"))
	}

	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)

	errStr := CheckSession(conn, token)
	if errStr != "" {
		return c.String(http.StatusUnauthorized, errStr)
	}

	var rooms []db.Room
	result := conn.Find(&rooms)
	if result.Error != nil {
		return c.String(http.StatusUnauthorized, ThrowError(result.Error.Error()+" (roomの検索エラー)"))
	}

	var roomsDetail []RoomDetail

	for _, v := range rooms {
		roomsDetail = append(roomsDetail, RoomDetail{
			ID:          strconv.Itoa(v.ID),
			Name:        v.Name,
			Description: v.Description,
			CreatedAt:   v.CreatedAt.String(),
		})
	}

	output := OutputGetRoomDetailList{
		Rooms: roomsDetail,
	}

	var res []byte
	res, err = json.Marshal(output)
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (JSONのMarshalエラー)"))
	}

	return c.String(http.StatusOK, string(res))
}

func CreateRoom(c echo.Context) error {
	p := new(InputCreateRoom)
	if err := c.Bind(p); err != nil {
		return c.String(http.StatusBadRequest, ThrowError(err.Error()+" (入力値エラー)"))
	}

	conn, err := db.InitDB()
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (DBの接続エラー)"))
	}

	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)

	errStr := CheckSession(conn, token)
	if errStr != "" {
		return c.String(http.StatusUnauthorized, errStr)
	}

	room := db.Room{
		Name:        p.Name,
		Description: p.Description,
	}

	result := conn.Create(&room)
	if result.Error != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (roomの作成エラー)"))
	}

	output := OutputCreateRoom{
		ID:          strconv.Itoa(room.ID),
		Name:        room.Name,
		Description: room.Description,
		CreatedAt:   room.CreatedAt.String(),
	}

	var res []byte
	res, err = json.Marshal(output)
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (JSONのMarshalエラー)"))
	}

	return c.String(http.StatusCreated, string(res))
}

func GetRoomDetail(c echo.Context) error {

	conn, err := db.InitDB()
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (DBの接続エラー)"))
	}

	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)

	errStr := CheckSession(conn, token)
	if errStr != "" {
		return c.String(http.StatusUnauthorized, errStr)
	}

	roomID := c.Param("roomId")

	var room db.Room
	result := conn.Find(&room, "id=?", roomID)
	if result.Error != nil {
		return c.String(http.StatusUnauthorized, ThrowError(result.Error.Error()+" (roomの検索エラー)"))
	}

	output := OutputGetRoomDetail{
		ID:          strconv.Itoa(room.ID),
		Name:        room.Name,
		Description: room.Description,
		CreatedAt:   room.CreatedAt.String(),
	}

	var res []byte
	res, err = json.Marshal(output)
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (JSONのMarshalエラー)"))
	}

	return c.String(http.StatusOK, string(res))
}

func CreateMessage(c echo.Context) error {
	conn, err := db.InitDB()
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (DBの接続エラー)"))
	}

	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)

	errStr := CheckSession(conn, token)
	if errStr != "" {
		return c.String(http.StatusUnauthorized, errStr)
	}

	p := new(InputCreateMessage)
	if err := c.Bind(p); err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (入力値エラー)"))
	}

	roomID, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (入力値エラー)"))
	}

	userID, err := strconv.Atoi(p.UserID)
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (入力値エラー)"))
	}

	message := db.Message{
		RoomID: roomID,
		UserID: userID,
		Text:   p.Text,
	}

	result := conn.Create(&message)
	if result.Error != nil {
		return c.String(http.StatusUnauthorized, ThrowError(result.Error.Error()+" (messageの作成エラー)"))
	}

	user := db.User{}

	result = conn.Find(&user, "id=?", message.UserID)
	if result.Error != nil {
		return c.String(http.StatusUnauthorized, ThrowError(result.Error.Error()+" (userの検索エラー)"))
	}

	output := OutputCreateMessage{
		ID:        strconv.Itoa(message.ID),
		Text:      message.Text,
		CreatedAt: message.CreatedAt.String(),
		User: User{
			ID:        strconv.Itoa(user.ID),
			Name:      user.Name,
			CreatedAt: user.CreatedAt.String(),
		},
	}

	var res []byte
	res, err = json.Marshal(output)
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (JSONのMarshalエラー)"))
	}

	return c.String(http.StatusCreated, string(res))
}

func CheckSession(conn *gorm.DB, token string) string {

	session := db.Session{}
	result := conn.First(&session, "token = ?", token)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return ThrowError(result.Error.Error() + " (sessionが見つからない)")
		}
		return ThrowError(result.Error.Error() + " (RecordNotFound以外のsessionの検索エラー)")
	}

	if session.ExpiredAt < time.Now().Unix() {
		session = db.Session{}
		conn.Delete(&session, "token = ?", token)
		return ThrowError("session expired" + " ログインし直してください")
	}

	return ""
}
