package appointment

import (
	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/MikelSot/metal-bat/model"
)

type Storage interface {
	GetTx() (model.Transaction, error)

	CreateTx(tx model.Transaction, m *model.Appointment) (model.Appointment, error)
	UpdateTx(tx model.Transaction, m *model.Appointment) error
	DeleteSoft(ID uint) error

	//TODO: debe traer todos las citas de hoy y que no esten canceladas, por taller
	GetAllAppointmentsDay(m model.Appointment) (model.Appointments, error)
	GetAllWhere(specification models.FieldsSpecification) (model.Appointment, error)
	GetWhere(specification models.FieldsSpecification) (model.Appointment, error)
}

type UseCase interface {
	CreateTx(tx model.Transaction, m *model.Appointment) (model.Appointment, error)
	UpdateTx(tx model.Transaction, m *model.Appointment) error
	DeleteSoft(ID uint) error

	GetAllAppointmentsDay(m model.Appointment) (model.Appointments, error)
	GetAllWhere(specification models.FieldsSpecification) (model.Appointment, error)
	GetWhere(specification models.FieldsSpecification) (model.Appointment, error)
}
