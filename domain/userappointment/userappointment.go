package userappointment

import (
	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/MikelSot/metal-bat/model"
)

//en este caso de uso es donde se hara toda la logica para crear una cita

type Storage interface {
	GetTx() (model.Transaction, error)
}

type UseCase interface {
	CreateTx(tx model.Transaction, m *model.Appointment) (model.Appointment, error)
	UpdateTx(tx model.Transaction, m *model.Appointment) error
	DeleteSoft(ID uint) error
}

type UseCaseAppointment interface {
	GetAllAppointmentsDay(m model.Appointment) (model.Appointments, error)
	GetAllWhere(specification models.FieldsSpecification) (model.Appointment, error)
	GetWhere(specification models.FieldsSpecification) (model.Appointment, error)
}
