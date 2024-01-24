package main

import (
	"fmt"
	"log"
	"os"

	"fizzBuzz/handlers"
	redisclient "fizzBuzz/redisClient"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	fmt.Println("Go Fizz-Buzz")

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
