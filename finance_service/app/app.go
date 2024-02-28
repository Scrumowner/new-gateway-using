package app

import (
	"context"
	"finance_service/internal/config"
	"finance_service/internal/modules/controller"
	"finance_service/internal/router"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	router     *router.Router
	controller *controller.Controller
	config     *config.Config
}

func NewApp(config *config.Config) *App {
	return &App{
		config:     config,
		controller: controller.NewController(config),
	}
}

func (a *App) Run() {
	router := router.NewRouter(a.controller)
	r := router.Router()
	serv := &http.Server{
		Addr:         fmt.Sprintf(":%s", a.config.Port),
		Handler:      r,
		ReadTimeout:  10,
		WriteTimeout: 10,
	}
	go func() {
		log.Printf(fmt.Sprintf("Server running on port %s", serv.Addr))
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("hutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
