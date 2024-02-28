package app

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"user_service/infrastructure/migrator"
	"user_service/internal"
	"user_service/internal/config"
	"user_service/internal/models"
	"user_service/internal/router"
)

type App struct {
	config      *config.Config
	controllers *internal.Controllers
	router      *router.Router
	migrator    *migrator.Migrator
}

func NewApp(db *sqlx.DB, cfg *config.Config) *App {
	controllers := internal.NewUserControllers(db)
	return &App{
		config:      cfg,
		controllers: controllers,
		router:      router.NewRouter(controllers),
		migrator:    migrator.NewMigrator(db),
	}

}

func (a *App) Run() {
	a.migrator.Migrate(models.User{})
	r := a.router.Route()
	server := http.Server{
		Addr:         a.config.Port,
		Handler:      r,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	go func() {
		log.Printf(fmt.Sprintf("Server running on port %s", server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("hutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
