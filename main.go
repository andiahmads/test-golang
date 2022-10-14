package main

import (
	"fmt"
	"log"
	"test-coding/game"
	"test-coding/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:endi@tcp(127.0.0.1:3306)/game-test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("connection to database good...")

	// call repository
	gameRepository := game.NewRepository(db)

	// call service
	gameService := game.NewService(gameRepository)

	// call controller
	gameHandler := handler.NewGameHandler(gameService)

	router := gin.Default()

	api := router.Group("/api")

	api.POST("/game", gameHandler.InsertData)
	api.GET("/game", gameHandler.ShowData)

	router.Run()

}
