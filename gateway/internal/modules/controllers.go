package modules

import (
	"gateway/internal/config"
	financecontorller "gateway/internal/modules/finance/controller"
	usercontroller "gateway/internal/modules/user/controller"
)

type Controllers struct {
	User *usercontroller.UserController
	Fin  *financecontorller.FinanceController
}

func NewControllers(cfg *config.Config) *Controllers {
	return &Controllers{
		User: usercontroller.NewUserController(cfg),
		Fin:  financecontorller.NewFinanceController(cfg),
	}
}
