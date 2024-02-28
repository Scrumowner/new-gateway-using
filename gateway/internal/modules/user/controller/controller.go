package controller

import (
	"bytes"
	"gateway/internal/config"
	"gateway/internal/infrastructure/responder"
	"gateway/internal/modules/user/provider"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type UserController struct {
	responder *responder.Responder
	provider  *provider.Provider
}

func NewUserController(cfg *config.Config) *UserController {
	return &UserController{
		responder: responder.NewResponder(),
		provider:  provider.NewProvider(cfg),
	}
}

func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {
	resp, err := u.provider.Login(r)
	if err != nil {
		u.responder.OutputInternalError(w, err)
	}
	buf := bytes.NewBuffer([]byte{})
	json.NewEncoder(buf).Encode(resp.Body)
	u.responder.OutputJson(w, buf.String())
}

func (u *UserController) Register(w http.ResponseWriter, r *http.Request) {
	resp, err := u.provider.Register(r)
	if err != nil {
		u.responder.OutputInternalError(w, err)
	}
	buf := bytes.NewBuffer([]byte{})
	json.NewEncoder(buf).Encode(resp.Body)
	u.responder.OutputJson(w, buf.String())
}
