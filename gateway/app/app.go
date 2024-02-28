package app

import (
	"context"
	"fmt"
	"gateway/internal/config"
	"gateway/internal/modules"
	"gateway/internal/router"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	config      *config.Config
	logger      *zap.SugaredLogger
	router      *router.MainRouter
	controllers *modules.Controllers
}

func NewApp() *App {
	zap := zap.Logger{}
	logger := zap.Sugar()
	cfg := config.NewConfig()
	return &App{
		config:      config.NewConfig(),
		logger:      logger,
		router:      router.NewRouter(),
		controllers: modules.NewControllers(cfg),
	}
}

func (a *App) Start() {
	a.config.Load()
	r := a.router.Route(a.controllers)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", a.config.Port),
		Handler: r,
	}
	go func() {
		log.Printf(fmt.Sprintf("Server running on port %s", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("hutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

}
