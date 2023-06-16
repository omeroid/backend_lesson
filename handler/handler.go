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
	//入力値の取得
	input := new(CreateUserInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error() + " (入力値エラー)",
		})
	}

	//DBのコネクションを取得
	conn := c.Get("db").(*gorm.DB)

	//パスワードのHash化処理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error() + " (hashしたpasswordの生成に失敗)",
		})
	}

	//usersテーブルにレコードを追加
	user := db.User{
		Name:         input.Username,
		PasswordHash: string(hashedPassword),
	}
	if result := conn.Create(&user); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: result.Error.Error() + " (user作成エラー)",
		})
	}

	output := CreateUserOutput{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
	}

	return c.JSON(http.StatusCreated, output)
}

// signin
func CheckUser(c echo.Context) error {
	//入力値の取得
	input := new(CheckUserInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error() + " (入力値エラー)",
		})
	}

	//DBのコネクションを取得
	conn := c.Get("db").(*gorm.DB)

	//usersテーブルにusernameで検索をかける
	user := db.User{}
	if result := conn.Take(&user, "name=?", input.Username); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: result.Error.Error() + " (user検索エラー)",
		})
	}

	//パスワードの比較(Hash化されている)
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error() + " (passwordが違う)",
		})
	}

	//sessionsテーブルに登録するtokenの生成（uuid）
	token, err := uuid.NewRandom()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error() + " (token生成エラー)",
		})
	}

	//Sessionへの追加
	session := db.Session{
		UserID:    user.ID,
		Token:     token.String(),
		ExpiredAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	if result := conn.Create(&session); result.Error != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: result.Error.Error() + " (session作成エラー)",
		})
	}

	output := CheckUserOutput{
		UserID:   user.ID,
		UserName: user.Name,
		Token:    token.String(),
	}

	return c.JSON(http.StatusOK, output)
}

// 全roomの情報取得
func GetRoomDetailList(c echo.Context) error {
	//DBのコネクションを取得
	conn := c.Get("db").(*gorm.DB)

	//sessionのtokenが有効か確認する
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error(),
		})
	}

	//roomsからレコードを全件取得
	var rooms []db.Room
	if result := conn.Find(&rooms); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: result.Error.Error() + " (room検索エラー)",
		})
	}

	//ユーザが必要なroomの情報を定義した構造体にデータを詰める
	var roomDetails []Room
	for _, v := range rooms {
		roomDetails = append(roomDetails, Room{
			ID:          v.ID,
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
	//入力値の取得
	input := new(CreateRoomInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error() + " (入力値エラー)",
		})
	}

	//DBのコネクションの取得
	conn := c.Get("db").(*gorm.DB)

	//sessionのtokenが有効か確認する
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error(),
		})
	}

	//roomsにレコードを挿入する
	room := db.Room{
		Name:        input.Name,
		Description: input.Description,
	}
	if result := conn.Create(&room); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: result.Error.Error() + " (room作成エラー)",
		})
	}

	output := CreateRoomOutput{
		ID:          room.ID,
		Name:        room.Name,
		Description: room.Description,
		CreatedAt:   room.CreatedAt.String(),
	}

	return c.JSON(http.StatusCreated, output)
}

// 指定したroomidのroomの詳細取得
func GetRoomDetail(c echo.Context) error {
	//DBのコネクションを取得する
	conn := c.Get("db").(*gorm.DB)

	//sessionのtokenが有効か確認する
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error(),
		})
	}

	//roomsをroomIDで検索する
	roomID := c.Param("roomId")
	var room db.Room
	if result := conn.Find(&room, "id=?", roomID); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: result.Error.Error() + " (room検索エラー)",
		})
	}

	output := GetRoomDetailOutput{
		ID:          room.ID,
		Name:        room.Name,
		Description: room.Description,
		CreatedAt:   room.CreatedAt.String(),
	}

	return c.JSON(http.StatusOK, output)
}

// messageをデータベースに登録する
func CreateMessage(c echo.Context) error {
	//DBのコネクションを取得
	conn := c.Get("db").(*gorm.DB)

	//sessionのtokenが有効か確認する
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error(),
		})
	}

	//リクエストのボディから入力値の取得
	input := new(CreateMessageInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error() + " (入力値エラー)",
		})
	}

	//リクエストのURLから入力値の取得
	roomID, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error() + " (roomID入力エラー)",
		})
	}

	//usersからuserIDで検索する
	user := db.User{}
	if result := conn.First(&user, "id=?", input.UserID); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: result.Error.Error() + " (user検索エラー)",
		})
	}

	//messagesにレコードを挿入する
	message := db.Message{
		RoomID: roomID,
		UserID: input.UserID,
		Text:   input.Text,
	}
	if result := conn.Create(&message); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: result.Error.Error() + " (message作成エラー)",
		})
	}

	output := CreateMessageOutput{
		ID:        message.ID,
		Text:      message.Text,
		CreatedAt: message.CreatedAt.String(),
		User: User{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt.String(),
		},
	}

	return c.JSON(http.StatusCreated, output)
}

// roomidで指定したroomのmessage詳細を全件取得
func GetMessageDetailList(c echo.Context) error {
	//DBのコネクションを取得
	conn := c.Get("db").(*gorm.DB)

	//sessionのtokenが有効か確認する
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error(),
		})
	}

	//リクエストのURLから入力値の取得
	roomID, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error() + " (roomID入力エラー)",
		})
	}

	//messagesにroomIDで検索をかける(一致全件取得)
	var messages []db.Message
	if result := conn.Find(&messages, "room_id=?", roomID); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: result.Error.Error() + "(message検索エラー)",
		})
	}

	//ユーザが必要な情報を定義した構造体にデータを詰める
	var messageDetails []Message
	for _, v := range messages {
		user := db.User{}
		conn.Find(&user, "id=?", v.UserID)
		messageDetails = append(messageDetails, Message{
			ID:        v.ID,
			Text:      v.Text,
			CreatedAt: v.CreatedAt.String(),
			User: User{
				ID:        user.ID,
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
	//DBのコネクションを取得する
	conn := c.Get("db").(*gorm.DB)

	//sessionのtokenが有効か確認する
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error(),
		})
	}

	//リクエストのURLから入力値の取得
	messageID := c.Param("messageId")
	roomID := c.Param("roomId")

	//messagesから指定したレコードを削除する
	message := &db.Message{}
	//deleteできてなかったときはmessageidに0がかえるのでそれで判定してほしい
	if result := conn.Clauses(clause.Returning{}).Where("id=? AND room_id=?", messageID, roomID).Delete(&message); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: result.Error.Error() + " (message削除エラー)",
		})
	}

	//usersからuserIDでuserを検索する
	user := db.User{}
	if result := conn.Find(&user, "id=?", message.UserID); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: result.Error.Error() + " (user検索エラー)",
		})
	}

	output := DeleteMessageOutput{
		ID:        message.ID,
		Text:      message.Text,
		CreatedAt: message.CreatedAt.String(),
		User: User{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt.String(),
		},
	}

	return c.JSON(http.StatusCreated, output)
}

// userにsessionが存在するか確認した後失効していないか確認する
func IsSessionValid(conn *gorm.DB, token string) error {
	//sessionsをtokenで検索する
	session := db.Session{}
	if result := conn.First(&session, "token = ?", token); result.Error != nil {
		return errors.New(result.Error.Error() + " (sessionの検索エラー)")
	}

	//tokenの有効期限が切れている時
	if session.ExpiredAt < time.Now().Unix() {
		session = db.Session{}
		conn.Delete(&session, "token = ?", token)
		return errors.New("session expired" + " ログインし直してください")
	}

	return nil
}
