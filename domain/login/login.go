package login

import "github.com/MikelSot/metal-bat/model"

type useCase interface {
	Login(m model.User) (model.User, error)
}

type UseCaseUser interface {
	ValidateUserPassword(email string) (model.User, error)
}
