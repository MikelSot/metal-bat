package login

import (
	"fmt"
	"strings"

	"github.com/MikelSot/metal-bat/model"
)

type Login struct {
	user UseCaseUser
}

func NewLogin(user UseCaseUser) Login {
	return Login{user}
}

func (l Login) Login(m model.User) (model.User, error) {
	e := model.NewError()

	if err := model.ValidateStructNil(m); err != nil {
		return model.User{}, fmt.Errorf("login: %w", err)
	}
	m.Email = strings.ToLower(m.Email)

	if !m.IsValidLenEmail() {
		e.SetError(fmt.Errorf("login: number of characters not allowed"))
		e.SetAPIMessage("¡Upps! número de caracteres del email no permitido")

		return model.User{}, e
	}

	if !m.IsEmailValidByRegex() {
		e.SetError(fmt.Errorf("login: invalid email"))
		e.SetAPIMessage("¡Upps! email no valido")

		return model.User{}, e
	}

	user, err := l.user.ValidateUserPassword(m.Email)
	if err != nil {
		e.SetError(fmt.Errorf("login: failed login attempt"))
		e.SetAPIMessage("¡Upps! intento de login fallido")

		return model.User{}, err
	}

	return user, nil
}
