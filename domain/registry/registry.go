package registry

import "github.com/MikelSot/metal-bat/model"

type Storage interface {
	GetTx() (model.Transaction, error)
}

type UseCase interface {
	Create(m *model.User) (model.User, error)
}

type UseCaseUser interface {
	CreateTx(tx model.Transaction, m *model.User) (model.User, error)

	GetByNickname(nickname string) (model.User, error)
	GetByEmail(email string) (model.User, error)
}
