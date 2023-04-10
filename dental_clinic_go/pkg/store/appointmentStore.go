package store

import "dental_clinic_go/internal/domain"

type AppointmentStore interface {
	GetByID(id int) (domain.Appointment, error)
	GetByDni(dni int) ([]domain.Appointment, error)
	Create(appointment domain.Appointment) (domain.Appointment, error)
	Update(appointment domain.Appointment) (bool, bool, domain.Appointment, error)
	Delete(id int) error
	CompleteEmptyAttributes(updatedAppointment domain.Appointment) (bool, bool, domain.Appointment, error)
}
