package controller

import (
	"finance_service/internal/config"
	"finance_service/internal/models"
	"finance_service/internal/modules/service"
	"finance_service/internal/responder"
	"fmt"
	"net/http"
)

type Controller struct {
	serivce   *service.Service
	responder *responder.Responder
}

func NewController(config *config.Config) *Controller {
	return &Controller{
		responder: responder.NewResponder(),
		serivce:   service.NewService(config),
	}
}

func (c *Controller) GetCoinst(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	req := &models.GetCoinsQuery{
		Page:       url.Get("page"),
		Limit:      url.Get("limit"),
		Currency:   url.Get("currency"),
		Blockcahin: url.Get("blockchain"),
	}
	json, err := c.serivce.GetCoins(req)
	if err != nil {
		c.responder.OutputInternalError(w, fmt.Errorf("Internal error))) "))
	}
	c.responder.OutputJson(w, json)
}
