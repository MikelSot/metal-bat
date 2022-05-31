package parameter

import (
	"database/sql"
	"github.com/MikelSot/metal-bat/model"
	"time"
)

type Parameter struct {
	db *sql.DB
}

func New(db *sql.DB) Parameter {
	return Parameter{db}
}

func (p Parameter) Create(m *model.Parameter) error {
	//TODO implement me
	panic("implement me")
}

func (p Parameter) Update(u uint, parameter *model.Parameter) error {
	//TODO implement me
	panic("implement me")
}

func (p Parameter) UpdateByName(name string, value string) error {
	//TODO implement me
	panic("implement me")
}

func (p Parameter) Delete(ID uint) error {
	//TODO implement me
	panic("implement me")
}

func (p Parameter) GetByID(u uint) (model.Parameter, error) {
	//TODO implement me
	panic("implement me")
}

func (p Parameter) GetByName(s string) (model.Parameter, error) {
	//TODO implement me
	panic("implement me")
}

func (p Parameter) GetAll() (model.Parameters, error) {
	//TODO implement me
	panic("implement me")
}

func (p Parameter) GetInt64(name string) (int64, bool) {
	//TODO implement me
	panic("implement me")
}

func (p Parameter) GetInt(name string) (int, bool) {
	//TODO implement me
	panic("implement me")
}

func (p Parameter) GetFloat32(name string) (float32, bool) {
	//TODO implement me
	panic("implement me")
}

func (p Parameter) GetTime(name string, format string) (time.Time, bool) {
	//TODO implement me
	panic("implement me")
}

func (p Parameter) GetBool(name string) (bool, bool) {
	//TODO implement me
	panic("implement me")
}
