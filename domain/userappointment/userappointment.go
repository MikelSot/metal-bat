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
	Create(m *model.Appointment) (model.Appointment, error)
	Update(m *model.Appointment) (model.Appointment, error)
	DeleteSoft(ID uint) error

	GetByUserID(userID uint) (model.Appointments, error)
}

type UseCaseAppointment interface {
	CreateTx(tx model.Transaction, m *model.Appointment) (model.Appointment, error)
	UpdateTx(tx model.Transaction, m *model.Appointment) (model.Appointment, error)
	DeleteSoft(ID uint) error

	GetAllAppointmentsDay() (model.Appointments, error)
	GetAllWhere(specification models.FieldsSpecification) (model.Appointment, error)
	GetWhere(specification models.FieldsSpecification) (model.Appointment, error)
}
