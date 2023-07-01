package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/omeroid/backend_lesson/backend/pkg/db"
	"github.com/omeroid/backend_lesson/backend/pkg/util"
	"gorm.io/gorm"
)

// 全roomの情報取得
func ListRoom(c echo.Context) error {
	//DBのコネクションを取得
	conn := c.Get("db").(*gorm.DB)

	//sessionのtokenが有効か確認する
	authHeader := c.Request().Header.Get("Authorization")
	token := util.ExtractBearerToken(authHeader)
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
	token := util.ExtractBearerToken(authHeader)
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
	token := util.ExtractBearerToken(authHeader)
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
