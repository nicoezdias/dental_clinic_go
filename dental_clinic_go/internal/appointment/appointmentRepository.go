package appointment

import (
	"dental_clinic_go/internal/domain"
	"dental_clinic_go/pkg/store"
	"errors"
	"fmt"
)

type AppointmentRepository interface {
	GetByID(id int) (domain.Appointment, error)
	GetByDni(dni int) ([]domain.Appointment, error)
	Create(a domain.Appointment) (domain.Appointment, error)
	CreateByDniAndLicense(dni int, license string, appointment domain.Appointment) (domain.Appointment, error)
	Update(id int, updatedAppointment domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}

type appointmentRepository struct {
	storage      store.AppointmentStore
	patientStore store.PatientStore
	dentistStore store.DentistStore
}

// NewAppointmentRepository crea un nuevo repositorio
func NewAppointmentRepository(storage store.AppointmentStore, patientStore store.PatientStore,
	dentistStore store.DentistStore) AppointmentRepository {
	return &appointmentRepository{storage, patientStore, dentistStore}
}

// GetByID busca un paciente por su id
func (r *appointmentRepository) GetByID(id int) (domain.Appointment, error) {
	appointment, err := r.storage.GetByID(id)
	if err != nil {
		return domain.Appointment{}, errors.New(fmt.Sprintf("appointment %d not found", id))
	}
	return appointment, nil
}

// GetByDni busca un paciente por su id
func (r *appointmentRepository) GetByDni(dni int) ([]domain.Appointment, error) {
	appointment, err := r.storage.GetByDni(dni)
	if err != nil {
		return []domain.Appointment{}, errors.New(fmt.Sprintf("appointments with patient.dni: %d not found", dni))
	}
	return appointment, nil
}

// Create agrega un nuevo paciente
func (r *appointmentRepository) Create(a domain.Appointment) (domain.Appointment, error) {
	appointment, err := r.storage.Create(a)
	if err != nil {
		return domain.Appointment{}, errors.New("error creating appointment")
	}
	return appointment, nil
}

// CreateByDniAndLicense agrega un nuevo turno por medio de el dni del paciente y la matricula del dentista
func (r *appointmentRepository) CreateByDniAndLicense(dni int, license string, appointment domain.Appointment) (domain.Appointment, error) {
	patient, err := r.patientStore.GetByDni(dni)
	if err != nil {
		return domain.Appointment{}, err
	}
	appointment.Patient = patient
	dentist, err := r.dentistStore.GetByLicense(license)
	if err != nil {
		return domain.Appointment{}, err
	}
	appointment.Dentist = dentist
	appointment, err = r.storage.Create(appointment)
	if err != nil {
		return domain.Appointment{}, errors.New("error creating appointment")
	}
	return appointment, nil
}

// Update actualiza un paciente
func (r *appointmentRepository) Update(id int, updatedAppointment domain.Appointment) (domain.Appointment, error) {
	updatedAppointment.Id = id
	patientFlag, dentistFlag, p, err := r.storage.Update(updatedAppointment)
	if err != nil {
		fmt.Println(err)
		return domain.Appointment{}, errors.New("error updating appointment")
	}
	if patientFlag {
		patient, err := r.patientStore.GetByID(p.Patient.Id)
		if err != nil {
			return domain.Appointment{}, err
		}
		p.Patient = patient
	}
	if dentistFlag {
		dentist, err := r.dentistStore.GetByID(p.Dentist.Id)
		if err != nil {
			return domain.Appointment{}, err
		}
		p.Dentist = dentist
	}
	return p, nil
}

// Delete busca un paciente por su id y lo elimina
func (r *appointmentRepository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
