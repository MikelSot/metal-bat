package user

import (
	"database/sql"

	"github.com/AJRDRGZ/db-query-builder/models"

	"github.com/MikelSot/metal-bat/infrastructure/postgres/tx"
	"github.com/MikelSot/metal-bat/model"
)

type User struct {
	db *sql.DB
	tx.Tx
}

func New(db *sql.DB) User {
	return User{db: db, Tx: tx.New(db)}
}

func (u User) CreateTx(tx model.Transaction, m *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u User) UpdateTx(tx model.Transaction, m *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u User) ResetPasswordTx(tx model.Transaction, m *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u User) UpdateNickname(m *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u User) DeleteSoft(ID uint) error {
	//TODO implement me
	panic("implement me")
}

func (u User) GetAllWhere(specification models.FieldsSpecification) (model.Users, error) {
	//TODO implement me
	panic("implement me")
}

func (u User) GetWhere(specification models.FieldsSpecification) (model.User, error) {
	//TODO implement me
	panic("implement me")
}
