package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/omeroid/kosen_backend_lesson/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"time"
)

// signup
func CreateUser(c echo.Context) error {
	p := new(InputCreateUser)
	if err := c.Bind(p); err != nil {
		return c.String(http.StatusBadRequest, ThrowError(err.Error()+" (入力値エラー)"))
	}

	conn := c.Get("db").(*gorm.DB)

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
	res, err := json.Marshal(output)
	if err != nil {
		return c.String(http.StatusBadRequest, ThrowError(err.Error()+" (jsonのMarshalのエラー)"))
	}

	return c.String(http.StatusCreated, string(res))

}

// signin
func CheckUser(c echo.Context) error {
	p := new(InputCheckUser)
	if err := c.Bind(p); err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (入力値エラー)"))
	}

	conn := c.Get("db").(*gorm.DB)

	user := db.User{}
	result := conn.Take(&user, "name=?", p.Username)
	if result.Error != nil {
		return c.String(http.StatusUnauthorized, ThrowError(result.Error.Error()+" (User検索エラー)"))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(p.Password)); err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (パスワードが違う)"))
	}

	token, err := uuid.NewRandom()
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (token生成エラー)"))
	}

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

// 全roomの情報取得
func GetRoomDetailList(c echo.Context) error {
	conn := c.Get("db").(*gorm.DB)

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

	var roomDetails []RoomDetail
	for _, v := range rooms {
		roomDetails = append(roomDetails, RoomDetail{
			ID:          strconv.Itoa(v.ID),
			Name:        v.Name,
			Description: v.Description,
			CreatedAt:   v.CreatedAt.String(),
		})
	}

	output := OutputGetRoomDetailList{
		Rooms: roomDetails,
	}

	var res []byte
	res, err := json.Marshal(output)
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (JSONのMarshalエラー)"))
	}

	return c.String(http.StatusOK, string(res))
}

// roomを作成する
func CreateRoom(c echo.Context) error {
	p := new(InputCreateRoom)
	if err := c.Bind(p); err != nil {
		return c.String(http.StatusBadRequest, ThrowError(err.Error()+" (入力値エラー)"))
	}

	conn := c.Get("db").(*gorm.DB)

	//Authorizationからtokenを取得してsessionの確認
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
		return c.String(http.StatusUnauthorized, ThrowError(result.Error.Error()+" (roomの作成エラー)"))
	}

	output := OutputCreateRoom{
		ID:          strconv.Itoa(room.ID),
		Name:        room.Name,
		Description: room.Description,
		CreatedAt:   room.CreatedAt.String(),
	}

	var res []byte
	res, err := json.Marshal(output)
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (JSONのMarshalエラー)"))
	}

	return c.String(http.StatusCreated, string(res))
}

// 指定したroomidのroomの詳細取得
func GetRoomDetail(c echo.Context) error {
	conn := c.Get("db").(*gorm.DB)

	//Authorizationからtokenを取得してsessionの確認
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
	res, err := json.Marshal(output)
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (JSONのMarshalエラー)"))
	}

	return c.String(http.StatusOK, string(res))
}

// messageをデータベースに登録する
func CreateMessage(c echo.Context) error {
	conn := c.Get("db").(*gorm.DB)

	//Authorizationからtokenを取得してsessionの確認
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

// roomidで指定したroomのmessage詳細を全件取得
func GetMessageDetailList(c echo.Context) error {
	conn := c.Get("db").(*gorm.DB)

	//Authorizationからtokenを取得してsessionの確認
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	errStr := CheckSession(conn, token)
	if errStr != "" {
		return c.String(http.StatusUnauthorized, errStr)
	}

	roomID, err := strconv.Atoi(c.Param("roomId"))

	var messages []db.Message

	result := conn.Find(&messages, "room_id=?", roomID)
	if result.Error != nil {
		return c.String(http.StatusUnauthorized, ThrowError(result.Error.Error()+" (messageの検索エラー)"))
	}

	var messageDetails []Message
	for _, v := range messages {
		user := db.User{}
		conn.Find(&user, "id=?", v.UserID)

		messageDetails = append(messageDetails, Message{
			ID:        strconv.Itoa(v.ID),
			Text:      v.Text,
			CreatedAt: v.CreatedAt.String(),
			User: User{
				ID:        strconv.Itoa(user.ID),
				Name:      user.Name,
				CreatedAt: user.CreatedAt.String(),
			},
		})
	}

	output := OutputGetMessageDetailList{
		Messages: messageDetails,
	}

	var res []byte

	res, err = json.Marshal(output)
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (JSONのMarshalエラー)"))
	}

	return c.String(http.StatusOK, string(res))
}

// messageをデータベースから削除
func DeleteMessage(c echo.Context) error {
	conn := c.Get("db").(*gorm.DB)

	//Authorizationからtokenを取得してsessionの確認
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	errStr := CheckSession(conn, token)
	if errStr != "" {
		return c.String(http.StatusUnauthorized, errStr)
	}

	messageID := c.Param("messageId")
	roomID := c.Param("roomId")

	message := &db.Message{}
	result := conn.Clauses(clause.Returning{}).Where("id=? AND room_id=?", messageID, roomID).Delete(&message)
	if result.Error != nil {
		return c.String(http.StatusUnauthorized, ThrowError(result.Error.Error()+" (messageの削除エラー)"))
	}
	//deleteできてなかったときはmessageidに"0"がかえるのでそれで判定してほしい

	user := db.User{}
	result = conn.Find(&user, "id=?", message.UserID)
	if result.Error != nil {
		return c.String(http.StatusUnauthorized, ThrowError(result.Error.Error()+" (userの検索エラー)"))
	}

	output := OutputDeleteMessage{
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

	res, err := json.Marshal(output)
	if err != nil {
		return c.String(http.StatusUnauthorized, ThrowError(err.Error()+" (JSONのMarshalエラー)"))
	}

	return c.String(http.StatusCreated, string(res))
}

// userにsessionが存在するか確認した後失効していないか確認する
func CheckSession(conn *gorm.DB, token string) string {

	session := db.Session{}
	result := conn.First(&session, "token = ?", token)
	if result.Error != nil {
		return ThrowError(result.Error.Error() + " (sessionの検索エラー)")
	}

	if session.ExpiredAt < time.Now().Unix() {
		session = db.Session{}
		conn.Delete(&session, "token = ?", token)
		return ThrowError("session expired" + " ログインし直してください")
	}

	return ""
}
