package main

import (
	"log"
	"os"

	"blog.phumaster.dev-api/database"
	"blog.phumaster.dev-api/model"
	"blog.phumaster.dev-api/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	conn, err := database.Initialize()
	database.Connection = conn
	if err != nil {
		panic(err)
	} else {
		conn.AutoMigrate(&model.User{}, &model.Article{})
		log.Println("Connection opened!")
	}
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	router.SetupRoutes(app)
	app.Logger.Fatal(app.Start(":" + port))
}
