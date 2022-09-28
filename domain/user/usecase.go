package user

import (
	"fmt"
	"strings"

	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/gosimple/slug"
	uuid "github.com/satori/go.uuid"

	"github.com/MikelSot/metal-bat/model"
)

var allowedFieldsForQuery = []string{
	"id", "nickname", "email", "is_active", "identification_type", "created_at", "updated_at", "deleted_at",
}

type User struct {
	storage Storage
}

func New(s Storage) User {
	return User{s}
}

func (u User) CreateTx(tx model.Transaction, m *model.User) (model.User, error) {
	e := model.NewError()

	if err := model.ValidateStructNil(m); err != nil {
		return model.User{}, fmt.Errorf("user: %w", err)
	}

	if u.hasAnEmptyField(*m) {
		e.SetError(fmt.Errorf("user: All fields are required"))
		e.SetAPIMessage("¡Upps! todos los campos son obligatorios")

		return model.User{}, e
	}

	if err := u.assignDefaultNickname(m); err != nil {
		return model.User{}, err
	}

	if err := u.storage.CreateTx(tx, m); err != nil {
		return model.User{}, fmt.Errorf("user.CreateTx(): %w", err)
	}

	return *m, nil
}

func (u User) UpdateTx(tx model.Transaction, m *model.User) (model.User, error) {
	if err := model.ValidateStructNil(m); err != nil {
		return model.User{}, fmt.Errorf("user.UpdateTx(): %w", err)
	}

	if !m.HasID() {
		return model.User{}, model.ErrInvalidID
	}

	if err := u.storage.UpdateTx(tx, m); err != nil {
		return model.User{}, fmt.Errorf("user.UpdateTx(): %w", err)
	}

	return *m, nil
}

func (u User) ResetPasswordTx(m *model.User) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("user.ResetPasswordTx(): %w", err)
	}

	if !m.HasID() {
		return model.ErrInvalidID
	}

	if !m.IsValidLenPassword() {
		e := model.NewError()
		e.SetError(fmt.Errorf("user.ResetPasswordTx()"))
		e.SetAPIMessage(fmt.Sprintf("¡Upps! La contraseña debe tener más de 6 dígitos"))

		return e
	}

	tx, err := u.storage.GetTx()
	if err != nil {
		return fmt.Errorf("user.storage.GetTx(): %w", err)
	}

	if err := u.storage.ResetPasswordTx(tx, m); err != nil {
		if errRollBack := tx.Rollback(); errRollBack != nil {
			return fmt.Errorf("user.ResetPasswordTx(): rollback error %s, %w", errRollBack.Error(), err)
		}

		return handleStorageErr(err)
	}

	return nil
}

func (u User) UpdateNickname(m *model.User) error {
	//TODO implement me
	//TODO: validar que el nuevo nickname sea unico

	panic("implement me")
}

func (u User) DeleteSoft(ID uint) error {
	if err := u.storage.DeleteSoft(ID); err != nil {
		return fmt.Errorf("user.DeleteSoft(): %w", err)
	}

	return nil
}

func (u User) ValidateUserPassword(email string) (model.User, error) {

	return model.User{}, nil
}

func (u User) GetByID(ID uint) (model.User, error) {
	return u.GetWhere(models.FieldsSpecification{
		Filters: models.Fields{{Name: "id", Value: ID}},
	})
}

func (u User) GetByNickname(nickname string) (model.User, error) {
	return u.GetWhere(models.FieldsSpecification{
		Filters: models.Fields{{Name: "nickname", Value: nickname}},
	})
}

func (u User) GetByEmail(email string) (model.User, error) {
	return u.GetWhere(models.FieldsSpecification{
		Filters: models.Fields{{Name: "email", Value: email}},
	})
}

func (u User) GetAllWhere(specification models.FieldsSpecification) (model.Users, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("user.GetAllWhere(): %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("user.GetAllWhere(): %w", err)
	}

	users, err := u.storage.GetAllWhere(specification)
	if err != nil {
		return nil, fmt.Errorf("user.GetAllWhere(): %w", err)
	}

	return users, nil
}

func (u User) GetWhere(specification models.FieldsSpecification) (model.User, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.User{}, fmt.Errorf("user.GetWhere(): %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.User{}, fmt.Errorf("user.GetWhere(): %w", err)
	}

	user, err := u.storage.GetWhere(specification)
	if err != nil {
		return model.User{}, fmt.Errorf("user.GetWhere(): %w", err)
	}

	return user, nil
}

func (u User) hasAnEmptyField(m model.User) bool {
	return m.IsStringEmpty(m.Firstname) || m.IsStringEmpty(m.Lastname) || m.IsStringEmpty(m.Email) || m.IsStringEmpty(m.Password)
}

func (u User) assignDefaultNickname(m *model.User) error {
	e := model.NewError()

	if m.IsStringEmpty(m.Nickname) {
		m.Nickname = m.Firstname + m.Lastname
	}

	m.Nickname = strings.ReplaceAll(m.Nickname, " ", "")
	m.Nickname = slug.Make(m.Nickname)

	user, err := u.GetByNickname(m.Nickname)
	if err != nil {
		return fmt.Errorf("user.u.GetByNickname(): %w", err)
	}
	if user.HasID() {
		e.SetError(fmt.Errorf("user: nickname already exists"))
		e.SetAPIMessage("¡Upps! el nickname ya existe")

		return e
	}

	random := uuid.NewV4().String()
	m.Nickname = fmt.Sprintf("%s%s", m.Nickname, random[:3])

	return nil
}

//TODO: eliminar los espacios al inicio y final de los nopmbres
//TODO: crear un nickname aleatorio
//TODO: crear una funciona IS VALID NICKNAME
//TODO: crear una funcion que em genere un nickname aletorio si es que se repite un nickname

// handleStorageErr handles errors from storage layer
func handleStorageErr(err error) error {
	e := model.NewError()
	e.SetError(err)

	switch err {
	default:
		return err
	}
}
