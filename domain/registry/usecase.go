package registry

import (
	"fmt"
	"github.com/MikelSot/metal-bat/model"
	"strings"
)

type Registry struct {
	storage Storage

	user UseCaseUser
}

func NewRegistry(storage Storage, user UseCaseUser) Registry {
	return Registry{
		storage,
		user,
	}
}

func (r Registry) Create(m *model.User) (model.User, error) {
	e := model.NewError()

	if err := model.ValidateStructNil(m); err != nil {
		return model.User{}, fmt.Errorf("registry: %w", err)
	}
	m.Email = strings.ToLower(m.Email)

	if !m.IsValidLenEmail() {
		e.SetError(fmt.Errorf("registry: number of characters not allowed"))
		e.SetAPIMessage("¡Upps! número de caracteres del email no permitido")

		return model.User{}, e
	}

	if !m.IsEmailValidByRegex() {
		e.SetError(fmt.Errorf("registry: invalid email"))
		e.SetAPIMessage("¡Upps! email no valido")

		return model.User{}, e
	}

	if !m.IsValidLenPassword() {
		e.SetError(fmt.Errorf("registry: very short password"))
		e.SetAPIMessage("¡Upps! contraseña muy corta")

		return model.User{}, e
	}

	user, err := r.user.GetByEmail(m.Email)
	if err != nil {
		return model.User{}, fmt.Errorf("registry.user.GetByEmail(): %w", err)
	}
	if user.HasID() {
		e.SetError(fmt.Errorf("registry: the email already exists"))
		e.SetAPIMessage("¡Upps! el correo ya existe")

		return model.User{}, e
	}

	tx, err := r.storage.GetTx()
	if err != nil {
		return model.User{}, fmt.Errorf("registry.storage.GetTx(): %w", err)
	}

	user, err = r.user.CreateTx(tx, m)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return model.User{}, fmt.Errorf("registry.tx.Rollback(): %w", errRollback)
		}

		return model.User{}, fmt.Errorf("registry.user.CreateTx(): %w", err)
	}

	return user, nil
}
