package controller

import (
	"bytes"
	"fmt"
	"github.com/jmoiron/sqlx"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"strings"
	"user_service/internal/models"
	"user_service/internal/modules/user/service"
	"user_service/internal/responder"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type UserController struct {
	service   *service.UserService
	responder *responder.Responder
}

func NewUserController(db *sqlx.DB) *UserController {
	return &UserController{
		service: service.NewUserSerivce(db),
	}
}
func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)
	token, err := u.service.GetUser(user)
	if err != nil {
		u.responder.OutputInternalError(w, err)
	}
	u.responder.OutputJson(w, string(token))
}

func (u *UserController) Register(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)
	err := u.service.SetUser(user)
	if err != nil {
		u.responder.OutputInternalError(w, err)
	}
	u.responder.OutputJson(w, string("Sucsecc register"))

}

func (u *UserController) Auth(w http.ResponseWriter, r *http.Request) {
	bearer := r.Header.Get("Authorization")
	s := strings.Split(bearer, " ")
	if len(s) != 2 {
		u.responder.OutputInternalError(w, fmt.Errorf("Invalid format of token"))
	}
	token := s[1]
	isAuth := u.service.Auth(token)
	if !isAuth {
		buf := bytes.NewBuffer([]byte{})
		_ = json.NewEncoder(buf).Encode(models.AuthResponse{
			IsAuth: false,
		})
		u.responder.OutputJson(w, buf.String())
	}
	buf := bytes.NewBuffer([]byte{})
	_ = json.NewEncoder(buf).Encode(models.AuthResponse{
		IsAuth: true,
	})
	u.responder.OutputJson(w, buf.String())

}
