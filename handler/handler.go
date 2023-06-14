package handler

import (
	"errors"
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
	input := new(CreateUserInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error() + " (入力値エラー)",
		})
	}

	conn := c.Get("db").(*gorm.DB)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10) //passwordをHash化する
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error() + " (hashしたpasswordの生成に失敗)",
		})
	}
	user := db.User{
		Name:         input.Username,
		PasswordHash: string(hashedPassword),
	}
	result := conn.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: result.Error.Error() + " (user作成エラー)",
		})
	}

	output := CreateUserOutput{
		ID:        strconv.Itoa(user.ID),
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
	}

	return c.JSON(http.StatusCreated, output)
}

// signin
func CheckUser(c echo.Context) error {
	input := new(CheckUserInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error() + " (入力値エラー)",
		})
	}

	conn := c.Get("db").(*gorm.DB)

	user := db.User{}
	result := conn.Take(&user, "name=?", input.Username)
	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: result.Error.Error() + " (user検索エラー)",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error() + " (passwordが違う)",
		})
	}

	token, err := uuid.NewRandom()
	if err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error() + " (token生成エラー)",
		})
	}

	//Sessionへの追加
	session := db.Session{
		UserID:    user.ID,
		Token:     token.String(),
		ExpiredAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	result = conn.Create(&session)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: result.Error.Error() + " (session作成エラー)",
		})
	}

	output := CheckUserOutput{
		UserID:   strconv.Itoa(user.ID),
		UserName: user.Name,
		Token:    token.String(),
	}

	return c.JSON(http.StatusOK, output)
}

// 全roomの情報取得
func GetRoomDetailList(c echo.Context) error {
	conn := c.Get("db").(*gorm.DB)

	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)

	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error(),
		})
	}

	var rooms []db.Room
	result := conn.Find(&rooms)
	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: result.Error.Error() + " (room検索エラー)",
		})
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

	output := GetRoomDetailListOutput{
		Rooms: roomDetails,
	}

	return c.JSON(http.StatusOK, output)
}

// roomを作成する
func CreateRoom(c echo.Context) error {
	input := new(CreateRoomInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error() + " (入力値エラー)",
		})
	}

	conn := c.Get("db").(*gorm.DB)

	//Authorizationからtokenを取得してsessionの確認
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error(),
		})
	}

	room := db.Room{
		Name:        input.Name,
		Description: input.Description,
	}

	result := conn.Create(&room)
	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: result.Error.Error() + " (room作成エラー)",
		})
	}

	output := CreateRoomOutput{
		ID:          strconv.Itoa(room.ID),
		Name:        room.Name,
		Description: room.Description,
		CreatedAt:   room.CreatedAt.String(),
	}

	return c.JSON(http.StatusCreated, output)
}

// 指定したroomidのroomの詳細取得
func GetRoomDetail(c echo.Context) error {
	conn := c.Get("db").(*gorm.DB)

	//Authorizationからtokenを取得してsessionの確認
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error(),
		})
	}

	roomID := c.Param("roomId")

	var room db.Room
	result := conn.Find(&room, "id=?", roomID)
	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: result.Error.Error() + " (room検索エラー)",
		})
	}

	output := GetRoomDetailOutput{
		ID:          strconv.Itoa(room.ID),
		Name:        room.Name,
		Description: room.Description,
		CreatedAt:   room.CreatedAt.String(),
	}

	return c.JSON(http.StatusOK, output)
}

// messageをデータベースに登録する
func CreateMessage(c echo.Context) error {
	conn := c.Get("db").(*gorm.DB)

	//Authorizationからtokenを取得してsessionの確認
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error(),
		})
	}

	input := new(CreateMessageInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error() + " (入力値エラー)",
		})
	}

	roomID, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error() + " (roomID入力エラー)",
		})
	}

	userID, err := strconv.Atoi(input.UserID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error() + " (userID入力エラー)",
		})
	}

	message := db.Message{
		RoomID: roomID,
		UserID: userID,
		Text:   input.Text,
	}

	result := conn.Create(&message)
	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: result.Error.Error() + " (message作成エラー)",
		})
	}

	user := db.User{}

	result = conn.Find(&user, "id=?", message.UserID)
	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: result.Error.Error() + " (user検索エラー)",
		})
	}

	output := CreateMessageOutput{
		ID:        strconv.Itoa(message.ID),
		Text:      message.Text,
		CreatedAt: message.CreatedAt.String(),
		User: User{
			ID:        strconv.Itoa(user.ID),
			Name:      user.Name,
			CreatedAt: user.CreatedAt.String(),
		},
	}

	return c.JSON(http.StatusCreated, output)
}

// roomidで指定したroomのmessage詳細を全件取得
func GetMessageDetailList(c echo.Context) error {
	conn := c.Get("db").(*gorm.DB)

	//Authorizationからtokenを取得してsessionの確認
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error(),
		})
	}

	roomID, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error() + " (roomID入力エラー)",
		})
	}

	var messages []db.Message

	result := conn.Find(&messages, "room_id=?", roomID)
	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: result.Error.Error() + "(message検索エラー)",
		})
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

	output := GetMessageDetailListOutput{
		Messages: messageDetails,
	}

	return c.JSON(http.StatusOK, output)
}

// messageをデータベースから削除
func DeleteMessage(c echo.Context) error {
	conn := c.Get("db").(*gorm.DB)

	//Authorizationからtokenを取得してsessionの確認
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error(),
		})
	}

	messageID := c.Param("messageId")
	roomID := c.Param("roomId")

	message := &db.Message{}
	result := conn.Clauses(clause.Returning{}).Where("id=? AND room_id=?", messageID, roomID).Delete(&message)
	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: result.Error.Error() + " (message削除エラー)",
		})
	}
	//deleteできてなかったときはmessageidに"0"がかえるのでそれで判定してほしい

	user := db.User{}
	result = conn.Find(&user, "id=?", message.UserID)
	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: result.Error.Error() + " (user検索エラー)",
		})
	}

	output := DeleteMessageOutput{
		ID:        strconv.Itoa(message.ID),
		Text:      message.Text,
		CreatedAt: message.CreatedAt.String(),
		User: User{
			ID:        strconv.Itoa(user.ID),
			Name:      user.Name,
			CreatedAt: user.CreatedAt.String(),
		},
	}

	return c.JSON(http.StatusCreated, output)
}

// userにsessionが存在するか確認した後失効していないか確認する
func IsSessionValid(conn *gorm.DB, token string) error {

	session := db.Session{}
	result := conn.First(&session, "token = ?", token)
	if result.Error != nil {
		return errors.New(result.Error.Error() + " (sessionの検索エラー)")
	}

	if session.ExpiredAt < time.Now().Unix() {
		session = db.Session{}
		conn.Delete(&session, "token = ?", token)
		return errors.New("session expired" + " ログインし直してください")
	}

	return nil
}
