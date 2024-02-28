package service

import (
	"finance_service/internal/config"
	"finance_service/internal/models"
	"finance_service/internal/modules/provider"
)

type Service struct {
	provider *provider.Provider
}

func NewService(config *config.Config) *Service {
	return &Service{
		provider: provider.NewProvider(config),
	}
}

func (s *Service) GetCoins(query *models.GetCoinsQuery) (string, error) {
	resp, err := s.provider.GetConis(
		query.Page,
		query.Limit,
		query.Currency,
		query.Blockcahin,
	)
	if err != nil {
		return "", err
	}
	return resp, nil
}
