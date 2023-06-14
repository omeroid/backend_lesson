package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/omeroid/kosen_backend_lesson/db"
	//"github.com/omeroid/kosen_backend_lesson/handler"
	"os"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//e.POST("/user/signup", handler.CreateUser)
	//e.POST("/user/signin", handler.CheckUser)
	//e.GET("/rooms", handler.GetRoomDetailList)
	//e.POST("/rooms", handler.CreateRoom)
	//e.GET("/rooms/:roomId", handler.GetRoomDetail)
	//e.POST("/rooms/:roomId/messages", handler.CreateMessage)
	//e.GET("/rooms/:roomId/messages", handler.GetMessageDetailList)
	//e.GET("/rooms/:roomId/messages/:messageId", handler.DeleteMessage)

	if err := godotenv.Load(".env"); err != nil {
		e.Logger.Fatal(".envの読み込み失敗: &v", err)
	}
	dbName := os.Getenv("DATABASE_NAME")

	conn, err := db.InitDB(dbName)
	if err != nil {
		e.Logger.Fatal("DBの接続失敗: &v", err)
	}
	db.Migrate(conn)
	db.InsertSampleRecord(conn)

	e.Logger.Fatal(e.Start(":1323"))
}
