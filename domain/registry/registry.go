package registry

import "github.com/MikelSot/metal-bat/model"

type UseCase interface {
	Create(m *model.User) (model.User, error)
}

type Storage interface {
	GetTx() (model.Transaction, error)
}

type UseCaseUser interface {
	CreateTx(tx model.Transaction, m *model.User) (model.User, error)

	GetByEmail(email string) (model.User, error)
}
