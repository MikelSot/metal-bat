package user

import (
	"github.com/AJRDRGZ/db-query-builder/models"

	"github.com/MikelSot/metal-bat/model"
)

type Storage interface {
	GetTx() (model.Transaction, error)

	CreateTx(tx model.Transaction, m *model.User) error
	UpdateTx(tx model.Transaction, m *model.User) error
	ResetPasswordTx(tx model.Transaction, m *model.User) error
	UpdateNickname(m *model.User) error
	DeleteSoft(ID uint) error

	GetAllWhere(specification models.FieldsSpecification) (model.Users, error)
	GetWhere(specification models.FieldsSpecification) (model.User, error)
}

type UseCase interface {
	CreateTx(tx model.Transaction, m *model.User) (model.User, error)
	UpdateTx(tx model.Transaction, m *model.User) (model.User, error)
	ResetPasswordTx(m *model.User) error
	UpdateNickname(m *model.User) error
	DeleteSoft(ID uint) error

	GetByID(ID uint) (model.User, error)
	GetByNickname(nickname string) (model.User, error)
	GetByEmail(email string) (model.User, error)
	GetAllWhere(specification models.FieldsSpecification) (model.Users, error)
	GetWhere(specification models.FieldsSpecification) (model.User, error)
}
