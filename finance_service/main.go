package main

import (
	app "finance_service/app"
	config "finance_service/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cfg := config.NewConfig()
	cfg.Load()
	app := app.NewApp(cfg)
	app.Run()

}
