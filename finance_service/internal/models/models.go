package models

type GetCoinsQuery struct {
	Page       string `json:"page"`
	Limit      string `json:"limit"`
	Currency   string `json:"currency"`
	Blockcahin string `json:"blockcahin"`
}
