package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/omeroid/kosen_backend_lesson/db"
	"github.com/omeroid/kosen_backend_lesson/handler"
)

// 現状ホームディレクトリ(Macなら"~"、 WindowsならC:\Users\ユーザ名)に.sqlitercというファイルを作りPRAGMA foreign_keys=ON;
// と記述しないと外部キー制約が効かない(解決策を考え中)
func main() {
	e := echo.New()

	//設定ファイルの読み込み
	if err := godotenv.Load(".env"); err != nil {
		e.Logger.Fatalf(".envの読み込み失敗: %s", err.Error())
	}
	dbName := os.Getenv("DATABASE_NAME")

	//DBへ接続
	conn, err := db.InitDB(dbName)
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

	//APIエンドポイントを定義する
	e.POST("/user/signup", handler.SignUp)
	e.POST("/user/signin", handler.SignIn)
	e.GET("/rooms", handler.ListRoom)
	e.POST("/rooms", handler.CreateRoom)
	e.GET("/rooms/:roomId", handler.GetRoom)
	e.POST("/rooms/:roomId/messages", handler.CreateMessage)
	e.GET("/rooms/:roomId/messages", handler.ListMessage)
	e.GET("/rooms/:roomId/messages/:messageId", handler.DeleteMessage)

	//localhost:1323でサーバ起動
	e.Logger.Fatal(e.Start(":1323"))
}
