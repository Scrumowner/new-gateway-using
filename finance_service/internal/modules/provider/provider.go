package provider

import (
	"finance_service/internal/config"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Provider struct {
	config *config.Config
}

func NewProvider(config *config.Config) *Provider {
	return &Provider{
		config: config,
	}
}

func (p *Provider) GetConis(page, limit, currency, blockchain string) (string, error) {
	url := fmt.Sprintf("https://openapiv1.coinstats.app/coins?page=%s&limit=%s&currency=%s&blockchain=%s", page, limit, currency, blockchain)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("accept", "application/json")
	key := p.config.Key
	request.Header.Add("X-API-KEY", key)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
