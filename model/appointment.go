package model

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

var (
	ErrAppointmentUserIDFK     = errors.New("appointment: The user(user_id) identification must be foreign")
	ErrAppointmentWorkshopIDFK = errors.New("appointment: The user(workshop_id) identification must be foreign")
	ErrAppointmentServiceIDFK  = errors.New("appointment: The user(service_id) identification must be foreign")
)

const (
	startServiceHours         = 8
	endServiceHours           = 19
	maximumAppointmentsPerDay = 20
)

type AppointmentStatus string

const (
	AppointmentStatusFailed    AppointmentStatus = "FAILED"
	AppointmentStatusInProcess AppointmentStatus = "IN_PROCESS"
	AppointmentStatusReview    AppointmentStatus = "REVIEW"
	AppointmentStatusDone      AppointmentStatus = "DONE"
)

type Appointment struct {
	ID              uint              `json:"id"`
	UserID          uint              `json:"user_id"`
	WorkshopID      uint16            `json:"workshop_id"`
	ServiceID       uint              `json:"service_id"`
	Description     string            `json:"description"`
	AppointmentDate time.Time         `json:"appointment_date"`
	DeliveryDate    time.Time         `json:"delivery_date"`
	AttentionOrder  uint16            `json:"attention_order"`
	State           AppointmentStatus `json:"state"`
	PickUp          bool              `json:"pick_up"`
	Delivery        bool              `json:"delivery"`
	IsCanceled      bool              `json:"is_canceled"`
	Slug            string            `json:"slug"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedAt       time.Time         `json:"deleted_at"`
}

func (a Appointment) HasID() bool { return a.ID > 0 }

func (a Appointment) IsStringEmpty(value string) bool { return len(value) == 0 }

func (a Appointment) Validate() error {
	newError := NewError()

	if a.AppointmentDate.IsZero() {
		newError.SetError(fmt.Errorf("Oops! Error appointment date must be empty"))
		newError.SetAPIMessage("¡Upps! Error la fecha de la cita debe estar vacio")

		return newError
	}

	if a.AppointmentDate.Before(time.Now()) {
		newError.SetError(fmt.Errorf("Oops! Error appointment date must be empty"))
		newError.SetAPIMessage("¡Upps! Error la fecha de la cita no debe ser una fecha anterior a la de hoy")

		return newError
	}

	if a.AppointmentDate.After(time.Now().AddDate(0, 0, 7)) {
		newError.SetError(fmt.Errorf("Oops! Error the appointment reservation must be made a maximum of one week in advance"))
		newError.SetAPIMessage("¡Upps! Error la reserva de cita se debe hacer como máximo con una semana de anticipación")

		return newError
	}

	if a.AppointmentDate.Hour() < startServiceHours || a.AppointmentDate.Hour() > endServiceHours {
		newError.SetError(fmt.Errorf("Oops! Error there is no service at the selected time"))
		newError.SetAPIMessage("¡Upps! Error no hay atención en la hora seleccionada")

		return newError
	}

	return nil
}

// Appointments  slice of Appointment
type Appointments []Appointment

func (a Appointments) IsEmpty() bool { return len(a) == 0 }

func (a Appointments) IsMaximumAppointmentsPerDay() bool { return len(a) == maximumAppointmentsPerDay }

func (a Appointments) GetAttentionOrder() uint16 {
	sort.Slice(a, func(i, j int) bool {
		return a[i].AttentionOrder < a[j].AttentionOrder
	})

	return a[len(a)-1].AttentionOrder + 1
}
