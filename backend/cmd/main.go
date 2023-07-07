package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/omeroid/backend_lesson/backend/handler"
	"github.com/omeroid/backend_lesson/backend/pkg/db"
	"github.com/omeroid/backend_lesson/backend/pkg/util"
)

func main() {
	e := echo.New()

	// databaseファイル
	databaseFilePath := util.JoinWithBackendRoot("chatapp.sqlite")
	//DBへ接続
	conn, err := db.InitDB(databaseFilePath)
	if err != nil {
		e.Logger.Fatalf("DBの接続失敗: %s", err.Error())
	}

	//テーブルを作成しサンプルデータを挿入
	if err := db.Migrate(conn); err != nil {
		e.Logger.Fatalf("DBへのmigration失敗: %s", err.Error())
	}
	if err := db.InsertSampleRecord(conn); err != nil {
		e.Logger.Fatalf("DBへのサンプルレコードのinsert失敗: %s", err.Error())
	}
	e.Logger.Print("Migration Successful")

	//middlewareを登録
	e.Use(db.DBMiddleware(conn)) //DBの接続をプールする
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	},
	))

	//APIエンドポイントを定義する
	e.POST("/user/signup", handler.SignUp)
	e.POST("/user/signin", handler.SignIn)
	e.GET("/rooms", handler.ListRoom)
	e.POST("/rooms", handler.CreateRoom)
	e.GET("/rooms/:roomId", handler.GetRoom)
	e.POST("/rooms/:roomId/messages", handler.CreateMessage)
	e.GET("/rooms/:roomId/messages", handler.ListMessage)
	e.DELETE("/rooms/:roomId/messages/:messageId", handler.DeleteMessage)

	//localhost:1323でサーバ起動
	e.Logger.Fatal(e.Start(":1323"))
}
