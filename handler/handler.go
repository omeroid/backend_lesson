package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/omeroid/kosen_backend_lesson/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// userを登録する
func SignUp(c echo.Context) error {
	//入力値の取得
	input := new(SignUpInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (入力値エラー)", err),
		})
	}

	//DBのコネクションを取得
	conn := c.Get("db").(*gorm.DB)

	if input.Username == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("(ユーザーネームが空)"),
		})
	}

	if input.Password == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("(パスワードが空)"),
		})
	}
	//パスワードのHash化処理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (hash化したpasswordの生成に失敗)", err),
		})
	}

	//usersテーブルにレコードを追加
	user := db.User{
		Name:         input.Username,
		PasswordHash: string(hashedPassword),
	}
	if result := conn.Create(&user); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (user作成エラー)", result.Error),
		})
	}

	output := SignUpOutput{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}

	return c.JSON(http.StatusCreated, output)
}

// ユーザを認証する
func SignIn(c echo.Context) error {
	//入力値の取得
	input := new(SignInInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (入力値エラー)", err),
		})
	}

	//DBのコネクションを取得
	conn := c.Get("db").(*gorm.DB)

	if input.Username == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("(ユーザーネームが空)"),
		})
	}

	if input.Password == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("(パスワードが空)"),
		})
	}

	//usersテーブルにusernameで検索をかける
	user := db.User{}
	if result := conn.Take(&user, "name=?", input.Username); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (user検索エラー)", result.Error),
		})
	}

	//過去のsessionの削除
	if result := conn.Where("user_id = ?", user.ID).Delete(&db.Session{}); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (過去session削除エラー)", result.Error),
		})
	}

	//パスワードの比較(Hash化されている)
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: fmt.Sprintf("%s (passwordが違う)", err),
		})
	}

	//sessionsテーブルに登録するtokenの生成（uuid）
	token, err := uuid.NewRandom()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: fmt.Sprintf("%s (token生成エラー)", err),
		})
	}

	//Sessionへの追加
	session := db.Session{
		UserID:    user.ID,
		Token:     token.String(),
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}
	if result := conn.Create(&session); result.Error != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (session作成エラー)", result.Error),
		})
	}

	output := SignInOutput{
		UserID:   user.ID,
		UserName: user.Name,
		Token:    token.String(),
	}

	return c.JSON(http.StatusOK, output)
}

// 全roomの情報取得
func ListRoom(c echo.Context) error {
	//DBのコネクションを取得
	conn := c.Get("db").(*gorm.DB)

	//sessionのtokenが有効か確認する
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: fmt.Sprintf("%s (tokenが無効)", err),
		})
	}

	//roomsからレコードを全件取得
	var rooms []db.Room
	if result := conn.Find(&rooms); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: fmt.Sprintf("%s (room検索エラー)", result.Error),
		})
	}

	//ユーザが必要なroomの情報を定義した構造体にデータを詰める
	var roomDetails []RoomOutput
	for _, v := range rooms {
		roomDetails = append(roomDetails, RoomOutput{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
		})
	}

	output := ListRoomOutput{
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
			Message: fmt.Sprintf("%s (入力値エラー)", err),
		})
	}

	//DBのコネクションの取得
	conn := c.Get("db").(*gorm.DB)

	//sessionのtokenが有効か確認する
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: fmt.Sprintf("%s (tokenが無効)", err),
		})
	}

	//roomsにレコードを挿入する
	room := db.Room{
		Name:        input.Name,
		Description: input.Description,
	}
	if result := conn.Create(&room); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (room作成エラー)", result.Error),
		})
	}

	output := CreateRoomOutput{
		ID:          room.ID,
		Name:        room.Name,
		Description: room.Description,
		CreatedAt:   room.CreatedAt,
	}

	return c.JSON(http.StatusCreated, output)
}

// 指定したroomidのroomの詳細取得
func GetRoom(c echo.Context) error {
	//DBのコネクションを取得する
	conn := c.Get("db").(*gorm.DB)

	//sessionのtokenが有効か確認する
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: fmt.Sprintf("%s (tokenが無効)", err),
		})
	}

	//roomsをroomIDで検索する
	roomID := c.Param("roomId")
	var room db.Room
	if result := conn.Find(&room, "id=?", roomID); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (room検索エラー)", result.Error),
		})
	}

	output := GetRoomOutput{
		ID:          room.ID,
		Name:        room.Name,
		Description: room.Description,
		CreatedAt:   room.CreatedAt,
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
			Message: fmt.Sprintf("%s (tokenが無効)", err),
		})
	}

	//リクエストのボディから入力値の取得
	input := new(CreateMessageInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (入力値エラー)", err),
		})
	}

	//リクエストのURLから入力値の取得
	roomID, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (roomID入力エラー)", err),
		})
	}

	//usersからuserIDで検索する
	user := db.User{}
	if result := conn.First(&user, "id=?", input.UserID); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (user検索エラー)", result.Error),
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
			Message: fmt.Sprintf("%s (message作成エラー)", result.Error),
		})
	}

	output := CreateMessageOutput{
		ID:        message.ID,
		Text:      message.Text,
		CreatedAt: message.CreatedAt,
		User: UserOutput{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
		},
	}

	return c.JSON(http.StatusCreated, output)
}

// roomidで指定したroomのmessage詳細を全件取得
func ListMessage(c echo.Context) error {
	//DBのコネクションを取得
	conn := c.Get("db").(*gorm.DB)

	//sessionのtokenが有効か確認する
	authHeader := c.Request().Header.Get("Authorization")
	token := ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: fmt.Sprintf("%s (tokenが無効)", err),
		})
	}

	//リクエストのURLから入力値の取得
	roomID, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (roomID入力エラー)", err),
		})
	}

	//messagesにroomIDで検索をかける(一致全件取得)
	var messages []db.Message
	if result := conn.Find(&messages, "room_id=?", roomID); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (message検索エラー)", result.Error),
		})
	}

	//ユーザが必要な情報を定義した構造体にデータを詰める
	var messageDetails []MessageOutput
	for _, v := range messages {
		user := db.User{}
		if result := conn.Find(&user, "id=?", v.UserID); result.Error != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Message: fmt.Sprintf("%s (user検索エラー)", result.Error),
			})
		}
		messageDetails = append(messageDetails, MessageOutput{
			ID:        v.ID,
			Text:      v.Text,
			CreatedAt: v.CreatedAt,
			User: UserOutput{
				ID:        user.ID,
				Name:      user.Name,
				CreatedAt: user.CreatedAt,
			},
		})
	}

	output := ListMessageOutput{
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
			Message: fmt.Sprintf("%s (tokenが無効)", err),
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
			Message: fmt.Sprintf("%s (message削除エラー)", result.Error),
		})
	}

	//usersからuserIDでuserを検索する
	user := db.User{}
	if result := conn.Find(&user, "id=?", message.UserID); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (user検索エラー)", result.Error),
		})
	}

	output := DeleteMessageOutput{
		ID:        message.ID,
		Text:      message.Text,
		CreatedAt: message.CreatedAt,
		User: UserOutput{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
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
	if session.ExpiredAt.Before(time.Now()) {
		return errors.New("session expired" + " ログインし直してください")
	}

	return nil
}
