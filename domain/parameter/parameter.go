package parameter

import (
	"github.com/MikelSot/metal-bat/model"
	"time"
)

type Storage interface {
	Create(m *model.Parameter) error
	Update(uint, *model.Parameter) error
	UpdateByName(name string, value string) error
	Delete(ID uint) error
	GetByID(uint) (model.Parameter, error)
	GetByName(string) (model.Parameter, error)
	GetAll() (model.Parameters, error)
}

type UseCase interface {
	Create(m *model.Parameter) error
	Update(uint, *model.Parameter) error
	UpdateByName(name string, value string) error
	Delete(ID uint) error
	GetByID(uint) (model.Parameter, error)
	GetByName(string) (string, error)
	GetAll() (model.Parameters, error)
	GetInt64(name string) (int64, bool)
	GetInt(name string) (int, bool)
	GetFloat32(name string) (float32, bool)
	GetTime(name string, format string) (time.Time, bool)
	GetBool(name string) (bool, bool)
}
