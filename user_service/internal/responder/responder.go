package responder

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Respond struct {
	message string `josn:"message"`
}
type Responder struct {
}

func NewResponder() *Responder {
	return &Responder{}
}

func (r *Responder) OutputJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	switch data.(type) {
	case string:
		fmt.Fprint(w, data)
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("responder json encode error")
	}
}

func (r *Responder) OutputUnautorized(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	respond := &Respond{
		message: err.Error(),
	}
	json.NewEncoder(w).Encode(&respond)

}

func (r *Responder) OutputInternalError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	respond := &Respond{
		message: err.Error(),
	}
	json.NewEncoder(w).Encode(&respond)
}
