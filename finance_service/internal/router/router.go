package router

import (
	"finance_service/internal/modules/controller"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	controller *controller.Controller
}

func NewRouter(controller *controller.Controller) *Router {
	return &Router{
		controller: controller,
	}

}

func (router *Router) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Get("/conis", router.controller.GetCoinst)
	})
	return r
}
