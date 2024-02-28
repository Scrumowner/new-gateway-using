package internal

import (
	"github.com/jmoiron/sqlx"
	"user_service/internal/modules/user/controller"
)

type Controllers struct {
	UserContorller *controller.UserController
}

func NewUserControllers(db *sqlx.DB) *Controllers {
	return &Controllers{
		UserContorller: controller.NewUserController(db),
	}
}
