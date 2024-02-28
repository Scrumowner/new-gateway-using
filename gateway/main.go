package main

import (
	app "gateway/app"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't load .env file")
	}
	app := app.NewApp()
	app.Start()
}
