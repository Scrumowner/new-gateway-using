package provider

import (
	"fmt"
	"gateway/internal/config"
	"log"
	"net/http"
)

type Provider struct {
	cfg *config.Config
}

func NewProvider(cfg *config.Config) *Provider {
	return &Provider{
		cfg: cfg,
	}
}
func (p *Provider) Login(r *http.Request) (*http.Response, error) {
	url := fmt.Sprintf("http://%s:%s/api/login", p.cfg.UserService.Host, p.cfg.UserService.Port)
	req, err := http.NewRequest("POST", url, r.Body)
	if err != nil {
		log.Println("Login route has invalid request")
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Login route has invalid request")
		return nil, err
	}
	return resp, nil
}

func (p *Provider) Register(r *http.Request) (*http.Response, error) {
	url := fmt.Sprintf("http://%s:%s/api/register", p.cfg.UserService.Host, p.cfg.UserService.Port)
	req, err := http.NewRequest("POST", url, r.Body)
	if err != nil {
		log.Println("Register route has invalid request")
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Login route has invalid request")
		return nil, err
	}
	return resp, nil
}

func (p *Provider) GetCurrency(r *http.Request) (*http.Response, error) {
	query := r.URL.Query()
	url := fmt.Sprintf("http://%s:%s/api/coins?page=%s&limit=%s&currency=%s&blockchain=%s",
		p.cfg.FinanceService.Host,
		p.cfg.FinanceService.Port,
		query.Get("page"),
		query.Get("limit"),
		query.Get("currency"),
		query.Get("blockchain"),
	)
	req, err := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Finance route has invalid request")
		return nil, err
	}
	return resp, nil
}
