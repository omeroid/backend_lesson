package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/omeroid/kosen_backend_lesson/db"
	"github.com/omeroid/kosen_backend_lesson/handler"
	"os"
)

// 現状ホームディレクトリ(Macなら"~"、 WindowsならC:\Users\ユーザ名)に.sqlitercというファイルを作りPRAGMA foreign_keys=ON;
// と記述しないと外部キー制約が効かない(解決策を考え中)
func main() {
	e := echo.New()

	//設定ファイルの読み込み
	if err := godotenv.Load(".env"); err != nil {
		e.Logger.Fatal(".envの読み込み失敗: &v", err)
	}
	dbName := os.Getenv("DATABASE_NAME")

	//DBへ接続
	conn, err := db.InitDB(dbName)
	if err != nil {
		e.Logger.Fatal("DBの接続失敗: &v", err)
	}

	//テーブルを作成しサンプルデータを挿入
	db.Migrate(conn)
	db.InsertSampleRecord(conn)
	fmt.Println("Migration Successful")

	//middlewareを登録
	e.Use(db.DBMiddleware(conn)) //DBの接続をプールする
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//APIエンドポイントを定義する
	e.POST("/user/signup", handler.CreateUser)
	e.POST("/user/signin", handler.CheckUser)
	e.GET("/rooms", handler.GetRoomDetailList)
	e.POST("/rooms", handler.CreateRoom)
	e.GET("/rooms/:roomId", handler.GetRoomDetail)
	e.POST("/rooms/:roomId/messages", handler.CreateMessage)
	e.GET("/rooms/:roomId/messages", handler.GetMessageDetailList)
	e.GET("/rooms/:roomId/messages/:messageId", handler.DeleteMessage)

	//localhost:1323でサーバ起動
	e.Logger.Fatal(e.Start(":1323"))
}
