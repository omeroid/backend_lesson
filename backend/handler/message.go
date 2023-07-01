package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/omeroid/backend_lesson/backend/pkg/db"
	"github.com/omeroid/backend_lesson/backend/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// messageをデータベースに登録する
func CreateMessage(c echo.Context) error {
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
	token := util.ExtractBearerToken(authHeader)
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
	token := util.ExtractBearerToken(authHeader)
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
