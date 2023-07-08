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

// メッセージをデータベースに登録する関数を定義します。
func CreateMessage(c echo.Context) error {
	//DBとの接続を確立します。
	conn := c.Get("db").(*gorm.DB)
	//セッションのtokenが有効かどうかを確認します。
	authHeader := c.Request().Header.Get("Authorization")
	token := util.ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		//Tokenが無効な場合、エラーレスポンスを返します。
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: fmt.Sprintf("%s (tokenが無効)", err),
		})
	}
	//リクエストボディから送られてきたデータ（メッセージ）を取得します。
	input := new(CreateMessageInput)
	if err := c.Bind(input); err != nil {
		//受け取りに問題があった場合、エラーレスポンスを返します。
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (入力値エラー)", err),
		})
	}
	//リクエストURLからルームIDを取得します。
	roomID, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		//ルームIDの取得に失敗した場合、エラーレスポンスを返します。
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (roomID入力エラー)", err),
		})
	}
	//指定されたユーザーIDに該当するユーザーをデータベースから検索します。
	user := db.User{}
	if result := conn.First(&user, "id=?", input.UserID); result.Error != nil {
		//ユーザーの検索に失敗した場合、エラーレスポンスを返します。
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (user検索エラー)", result.Error),
		})
	}
	//新規メッセージをデータベースに登録します。
	message := db.Message{
		RoomID: roomID,
		UserID: input.UserID,
		Text:   input.Text,
	}
	if result := conn.Create(&message); result.Error != nil {
		//メッセージの作成に失敗した場合、エラーレスポンスを返します。
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (message作成エラー)", result.Error),
		})
	}
	//返すべきメッセージデータを作成します。
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
	//作成したメッセージデータをJSONとして返します。
	return c.JSON(http.StatusCreated, output)
}

// 特定のroomでの全てのmessageを取得する関数を定義します。
func ListMessage(c echo.Context) error {
	//DBとの接続を確立します。
	conn := c.Get("db").(*gorm.DB)
	//セッションのtokenが有効かどうかを確認します。
	authHeader := c.Request().Header.Get("Authorization")
	token := util.ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		//Tokenが無効な場合、エラーレスポンスを返します。
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: fmt.Sprintf("%s (tokenが無効)", err),
		})
	}
	//リクエストURLからルームIDを取得します。
	roomID, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		//ルームIDの取得に失敗した場合、エラーレスポンスを返します。
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (roomID入力エラー)", err),
		})
	}
	//指定されたルームIDを持つメッセージをデータベースから取得します。
	var messages []db.Message
	if result := conn.Find(&messages, "room_id=?", roomID); result.Error != nil {
		//メッセージの取得に失敗した場合、エラーレスポンスを返します。
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (message検索エラー)", result.Error),
		})
	}
	//取得したメッセージとそれに紐づくユーザー情報を整形します。
	//最終的にはユーザーに表示するために必要なメッセージの詳細を作成します。
	var messageDetails []MessageOutput
	for _, v := range messages {
		user := db.User{}
		if result := conn.Find(&user, "id=?", v.UserID); result.Error != nil {
			//ユーザーの検索に失敗した場合、エラーレスポンスを返します。
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
	//取得したメッセージの詳細をJSONとして返します。
	output := ListMessageOutput{
		Messages: messageDetails,
	}
	return c.JSON(http.StatusOK, output)
}

// 特定のメッセージをデータベースから削除する関数を定義します。
func DeleteMessage(c echo.Context) error {
	//DBとの接続を確立します。
	conn := c.Get("db").(*gorm.DB)
	//セッションのtokenが有効かどうかを確認します。
	authHeader := c.Request().Header.Get("Authorization")
	token := util.ExtractBearerToken(authHeader)
	if err := IsSessionValid(conn, token); err != nil {
		//Tokenが無効な場合、エラーレスポンスを返します。
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: fmt.Sprintf("%s (tokenが無効)", err),
		})
	}
	//リクエストURLからメッセージIDとルームIDを取得します。
	messageID := c.Param("messageId")
	roomID := c.Param("roomId")
	//指定されたメッセージIDとルームIDに該当するメッセージをデータベースから削除します。
	message := &db.Message{}
	if result := conn.Clauses(clause.Returning{}).Where("id=? AND room_id=?", messageID, roomID).Delete(&message); result.Error != nil {
		//メッセージの削除に失敗した場合、エラーレスポンスを返します。
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (message削除エラー)", result.Error),
		})
	}
	//削除したメッセージと関連するユーザーをデータベースから検索します。
	user := db.User{}
	if result := conn.Find(&user, "id=?", message.UserID); result.Error != nil {
		//ユーザーの検索に失敗した場合、エラーレスポンスを返します。
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (user検索エラー)", result.Error),
		})
	}
	//削除したメッセージの詳細を作成します。
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
	//削除したメッセージの詳細をJSONとして返します。
	return c.JSON(http.StatusCreated, output)
}
