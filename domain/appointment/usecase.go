package appointment

import (
	"fmt"

	"github.com/AJRDRGZ/db-query-builder/models"

	"github.com/MikelSot/metal-bat/model"
)

var allowedFieldsForQuery = []string{"id"}

type Appointment struct {
	storage Storage
}

func New(s Storage) Appointment {
	return Appointment{s}
}

func (a Appointment) CreateTx(tx model.Transaction, m *model.Appointment) (model.Appointment, error) {
	if err := model.ValidateStructNil(m); err != nil {
		return model.Appointment{}, fmt.Errorf("appointment: %w", err)
	}

	if err := a.storage.CreateTx(tx, m); err != nil {
		return model.Appointment{}, fmt.Errorf("appointment.CreateTx(): %w", err)
	}

	return *m, nil
}

func (a Appointment) UpdateTx(tx model.Transaction, m *model.Appointment) (model.Appointment, error) {
	if err := model.ValidateStructNil(m); err != nil {
		return model.Appointment{}, fmt.Errorf("appointment: %w", err)
	}

	if !m.HasID() {
		return model.Appointment{}, model.ErrInvalidID
	}

	if err := a.storage.UpdateTx(tx, m); err != nil {
		return model.Appointment{}, fmt.Errorf("appointment.UpdateTx(): %w", err)
	}

	return *m, nil
}

func (a Appointment) DeleteSoft(ID uint) error {
	if err := a.storage.DeleteSoft(ID); err != nil {
		return fmt.Errorf("appointment.DeleteSoft(): %w", err)
	}

	return nil
}

func (a Appointment) GetAllAppointmentsDay() (model.Appointments, error) {
	//TODO implement me
	panic("implement me")
}

func (a Appointment) GetAllWhere(specification models.FieldsSpecification) (model.Appointments, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("appointment.GetAllWhere(): %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("appointment.GetAllWhere(): %w", err)
	}

	appointments, err := a.storage.GetAllWhere(specification)
	if err != nil {
		return nil, fmt.Errorf("appointment.GetAllWhere(): %w", err)
	}

	return appointments, nil
}

func (a Appointment) GetWhere(specification models.FieldsSpecification) (model.Appointment, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.Appointment{}, fmt.Errorf("appointment.GetWhere(): %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.Appointment{}, fmt.Errorf("appointment.GetWhere(): %w", err)
	}

	appointment, err := a.storage.GetWhere(specification)
	if err != nil {
		return model.Appointment{}, fmt.Errorf("appointment.GetWhere(): %w", err)
	}

	return appointment, nil
}
