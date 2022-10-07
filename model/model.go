package model

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrNilPointer        = errors.New("el modelo recibido es nulo")
	ErrInvalidID         = errors.New("el id recibido no es valido")
	ErrModelNotFound     = errors.New("el modelo no fue encontrado")
	ErrParameterNotFound = errors.New("el parámetro no se encuentra configurado")
)

func ValidateStructNil(i interface{}) error {
	if reflect.ValueOf(i).Kind() == reflect.Struct {
		return nil
	}

	if i == nil {
		return ErrNilPointer
	}

	if reflect.ValueOf(i).IsNil() {
		return ErrNilPointer
	}

	return nil
}

func errRequiredField(field string) error {
	e := NewError()
	e.SetError(fmt.Errorf("missing %s field", field))
	e.SetAPIMessage(fmt.Sprintf("¡Upps! no enviaste el campo: %s", field))
	return e
}
