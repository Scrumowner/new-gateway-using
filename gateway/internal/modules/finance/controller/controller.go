package controller

import (
	"bytes"
	"encoding/json"
	"gateway/internal/config"
	"gateway/internal/infrastructure/responder"
	"gateway/internal/modules/user/provider"
	"net/http"
)

type FinanceController struct {
	responder *responder.Responder
	provider  *provider.Provider
}

func NewFinanceController(cfg *config.Config) *FinanceController {
	return &FinanceController{
		responder: responder.NewResponder(),
		provider:  provider.NewProvider(cfg),
	}
}

func (f *FinanceController) GetCurrency(w http.ResponseWriter, r *http.Request) {
	resp, err := f.provider.GetCurrency(r)
	if err != nil {
		f.responder.OutputInternalError(w, err)
	}
	buf := bytes.NewBuffer([]byte{})
	json.NewEncoder(buf).Encode(resp.Body)
	f.responder.OutputJson(w, buf.String())
}
