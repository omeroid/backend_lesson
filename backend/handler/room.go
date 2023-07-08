package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/omeroid/backend_lesson/backend/pkg/db"
	"github.com/omeroid/backend_lesson/backend/pkg/util"
	"gorm.io/gorm"
)

// 全roomの情報取得関数
func ListRoom(c echo.Context) error {
	//DBのコネクションを取得
	conn := c.Get("db").(*gorm.DB)

	//リクエストヘッダーからトークンを取得してバリデーションをする
	authHeader := c.Request().Header.Get("Authorization")
	token := util.ExtractBearerToken(authHeader)
	//セッションが無効ならエラーメッセージを返す
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: fmt.Sprintf("%s (tokenが無効)", err),
		})
	}

	//全てのroomデータをデータベースから取得
	var rooms []db.Room
	//取得に失敗した場合はエラーメッセージを返す
	if result := conn.Find(&rooms); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: fmt.Sprintf("%s (room検索エラー)", result.Error),
		})
	}

	//取得したroomデータを構造体に格納
	var roomDetails []RoomOutput
	for _, v := range rooms {
		roomDetails = append(roomDetails, RoomOutput{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
		})
	}

	//レスポンス用のデータを作成
	output := ListRoomOutput{
		Rooms: roomDetails,
	}

	//レスポンスデータをJSON形式で返す
	return c.JSON(http.StatusOK, output)
}

// roomを作成する関数
func CreateRoom(c echo.Context) error {
	//クライアントから送られてきたデータを取得
	input := new(CreateRoomInput)
	//データのバリデーションを行い、エラーが発生したらエラーメッセージを返す
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (入力値エラー)", err),
		})
	}

	//DBのコネクションの取得
	conn := c.Get("db").(*gorm.DB)

	//リクエストヘッダーからトークンを取得してバリデーションをする
	authHeader := c.Request().Header.Get("Authorization")
	token := util.ExtractBearerToken(authHeader)
	//セッションが無効ならエラーメッセージを返す
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: fmt.Sprintf("%s (tokenが無効)", err),
		})
	}

	//新規Roomデータを作成
	room := db.Room{
		Name:        input.Name,
		Description: input.Description,
	}
	//RoomデータをDBに保存して、失敗したらエラーメッセージを返す
	if result := conn.Create(&room); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (room作成エラー)", result.Error),
		})
	}

	// 作成されたRoomのレスポンスデータを作成する
	output := CreateRoomOutput{
		ID:          room.ID,
		Name:        room.Name,
		Description: room.Description,
		CreatedAt:   room.CreatedAt,
	}

	//レスポンスデータをJSON形式で返す
	return c.JSON(http.StatusCreated, output)
}

// 指定したroomidのroomの詳細取得関数
func GetRoom(c echo.Context) error {
	//DBのコネクションを取得する
	conn := c.Get("db").(*gorm.DB)

	//リクエストヘッダーからトークンを取得してバリデーションをする
	authHeader := c.Request().Header.Get("Authorization")
	token := util.ExtractBearerToken(authHeader)
	//セッションが無効ならエラーメッセージを返す
	if err := IsSessionValid(conn, token); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: fmt.Sprintf("%s (tokenが無効)", err),
		})
	}

	//パラメータからroomIdを取得して該当roomをDBから検索
	roomID := c.Param("roomId")
	var room db.Room
	//検索に失敗した場合はエラーメッセージを返す
	if result := conn.Find(&room, "id=?", roomID); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (room検索エラー)", result.Error),
		})
	}

	//取得したroomのレスポンスデータを作成する
	output := GetRoomOutput{
		ID:          room.ID,
		Name:        room.Name,
		Description: room.Description,
		CreatedAt:   room.CreatedAt,
	}

	//レスポンスデータをJSON形式で返す
	return c.JSON(http.StatusOK, output)
}
