package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"training.go/fizzBuzz/handlers"
	redisclient "training.go/fizzBuzz/redisClient"
)

func main() {

	fmt.Println("Go Fizz-Buzz")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env")
	}

	client, err := redisclient.StartRedis()
	if err != nil {
		log.Printf("Erreur : %v", err)
		os.Exit(1)
	}
	defer client.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/fizzbuzz", func(c echo.Context) error {
		return handlers.FizzBuzzHandler(c, client)
	})

	e.GET("/statistics", func(c echo.Context) error {
		return handlers.StatisticsHandler(c, client)
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "9090"
	}
	e.Logger.Fatal(e.Start(":" + httpPort))
}
