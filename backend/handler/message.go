package handler

import (
	// "fmt"
	// "net/http"
	// "strconv"

	"github.com/labstack/echo/v4"
	// "github.com/omeroid/backend_lesson/backend/pkg/db"
	// "github.com/omeroid/backend_lesson/backend/pkg/util"
	// "gorm.io/gorm"
)

// TODO A: メッセージ作成機能を実装してください
func CreateMessage(c echo.Context) error {
	// 1. データベース接続を取得（c.Get("db")を*gorm.DBに型変換）
	// 2. Authorizationヘッダーからトークンを取得し、util.ExtractBearerTokenで抽出
	// 3. IsSessionValidでトークンを検証
	// 4. リクエストボディからCreateMessageInputを取得
	// 5. URLパラメータからroomIdを取得し、strconv.Atoiで整数に変換
	// 6. UserIDに該当するユーザーをデータベースから検索
	// 7. メッセージをデータベースに保存
	// 8. CreateMessageOutputを作成して返却（http.StatusCreated）
	return nil
}

// TODO B: メッセージ一覧取得機能を実装してください
func ListMessage(c echo.Context) error {
	// 1. データベース接続を取得（c.Get("db")を*gorm.DBに型変換）
	// 2. Authorizationヘッダーからトークンを取得し、util.ExtractBearerTokenで抽出
	// 3. IsSessionValidでトークンを検証
	// 4. URLパラメータからroomIdを取得し、strconv.Atoiで整数に変換
	// 5. 指定されたルームIDのメッセージをデータベースから取得
	// 6. 各メッセージにユーザー情報を追加してListMessageOutputを作成
	// 7. レスポンスを返却（http.StatusOK）
	return nil
}

// TODO C: メッセージ削除機能を実装してください
func DeleteMessage(c echo.Context) error {
	// 1. データベース接続を取得（c.Get("db")を*gorm.DBに型変換）
	// 2. Authorizationヘッダーからトークンを取得し、util.ExtractBearerTokenで抽出
	// 3. IsSessionValidでトークンを検証
	// 4. URLパラメータからmessageIdとroomIdを取得し、strconv.Atoiで整数に変換
	// 5. 指定されたメッセージをデータベースから削除
	// 6. レスポンスを返却（http.StatusNoContent）
	return nil
}
