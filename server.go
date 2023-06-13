package main

import (
	//"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/omeroid/kosen_backend_lesson/db"
	"github.com/omeroid/kosen_backend_lesson/handler"
)

func main() {
	conn, _ := db.InitDB()
	db.Migrate(conn)
	db.Insert(conn)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//e.GET("/room/:id", handler.GetMessages)
	//e.POST("/lobby", handler.SendMessage) //roomIDとmsgが必要
	//e.POST("/delete", handler.DeleteMessage)
	e.POST("/user/signup", handler.CreateUser)
	e.POST("/user/signin", handler.CheckUser)
	e.GET("/rooms", handler.GetRoomDetailList)
	e.POST("/rooms", handler.CreateRoom)
	e.GET("/rooms/:roomId", handler.GetRoomDetail)

	e.Logger.Fatal(e.Start(":1323"))
}
