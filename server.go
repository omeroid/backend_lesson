package main

import (
	//	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/omeroid/kosen_backend_lesson/handler"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/room/:id", handler.GetMessages)
	e.POST("/lobby", handler.SendMessage) //roomIDとmsgが必要

	e.Logger.Fatal(e.Start(":1323"))
}
