package router

import (
	"github.com/go-chi/chi/v5"
	"user_service/internal"
)

type Router struct {
	Controllers *internal.Controllers
}

func NewRouter(controllers *internal.Controllers) *Router {
	return &Router{
		Controllers: controllers,
	}
}

func (r *Router) Route() *chi.Mux {
	router := chi.NewRouter()
	router.Route("/user", func(router chi.Router) {
		router.Post("/login", r.Controllers.UserContorller.Register)
		router.Post("/register", r.Controllers.UserContorller.Login)
		router.Post("/auth", r.Controllers.UserContorller.Auth)
	})
	return router
}
