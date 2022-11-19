package userappointment

import (
	"fmt"

	"github.com/MikelSot/metal-bat/model"
)

type UserAppointment struct {
	storage     Storage
	appointment UseCaseAppointment
}

func New(s Storage, a UseCaseAppointment) UserAppointment {
	return UserAppointment{s, a}
}

func (u UserAppointment) Create(m *model.Appointment) (model.Appointment, error) {
	if err := model.ValidateStructNil(m); err != nil {
		return model.Appointment{}, fmt.Errorf("userappointment: %w", err)
	}

	if err := m.Validate(); err != nil {
		return model.Appointment{}, err
	}

	newError := model.NewError()
	appointmentsPerDay, err := u.appointment.GetAllAppointmentsDay()
	if err != nil {
		return model.Appointment{}, fmt.Errorf("userappointment.appointment.GetAllAppointmentsDay(): %w", err)
	}

	if appointmentsPerDay.IsMaximumAppointmentsPerDay() {
		newError.SetError(fmt.Errorf("Oops! Error the limit of appointments per day was reached"))
		newError.SetAPIMessage("¡Upps! Error se alcanzo el limite de citas por dia")

		return model.Appointment{}, newError
	}

	m.State = model.AppointmentStatusInProcess
	m.AttentionOrder = appointmentsPerDay.GetAttentionOrder()

	tx, err := u.storage.GetTx()
	if err != nil {
		return model.Appointment{}, fmt.Errorf("userappointment.storage.GetTx(): %w", err)
	}

	appointment, err := u.appointment.CreateTx(tx, m)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return model.Appointment{}, fmt.Errorf("userappointment:  %s, %w", errRollback, err)
		}

		return model.Appointment{}, fmt.Errorf("userappointment: %w", err)
	}

	return appointment, nil
}

func (u UserAppointment) Update(m *model.Appointment) (model.Appointment, error) {
	if err := model.ValidateStructNil(m); err != nil {
		return model.Appointment{}, fmt.Errorf("userappointment: %w", err)
	}

	if err := m.Validate(); err != nil {
		return model.Appointment{}, err
	}

	tx, err := u.storage.GetTx()
	if err != nil {
		return model.Appointment{}, fmt.Errorf("userappointment.storage.GetTx(): %w", err)
	}

	appointment, err := u.appointment.CreateTx(tx, m)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return model.Appointment{}, fmt.Errorf("userappointment:  %s, %w", errRollback, err)
		}

		return model.Appointment{}, fmt.Errorf("userappointment: %w", err)
	}

	return appointment, nil
}

func (u UserAppointment) DeleteSoft(ID uint) error {
	//TODO implement me
	panic("implement me")
}

func (u UserAppointment) GetByUserID(userID uint) (model.Appointments, error) {
	//TODO implement me
	panic("implement me")
}
