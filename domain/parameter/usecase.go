package parameter

import (
	"fmt"
	"github.com/MikelSot/metal-bat/model"
	"time"
)

type Parameter struct {
	storage Storage
}

func New(storage Storage) Parameter {
	return Parameter{storage}
}

func (p Parameter) Create(parameter *model.Parameter) error {
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

func (p Parameter) GetByName(s string) (string, error) {
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
	fmt.Println(name)
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
