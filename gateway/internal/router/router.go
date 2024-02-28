package router

import (
	"gateway/internal/modules"
	"github.com/go-chi/chi/v5"
)

type MainRouter struct {
}

func NewRouter() *MainRouter {
	return &MainRouter{}
}

func (m *MainRouter) Route(controllers *modules.Controllers) *chi.Mux {
	r := chi.NewRouter()
	r.Route("/user", func(r chi.Router) {
		r.Post("/login", controllers.User.Login)
		r.Post("/register", controllers.User.Register)
	})
	r.Route("/finance", func(r chi.Router) {
		r.Get("/conis", controllers.Fin.GetCurrency)
	})
	return r
}
