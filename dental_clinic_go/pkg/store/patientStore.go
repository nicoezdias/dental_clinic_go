package store

import "dental_clinic_go/internal/domain"

type PatientStore interface {
	GetByID(id int) (domain.Patient, error)
	GetByDni(dni int) (domain.Patient, error)
	Create(patient domain.Patient) (domain.Patient, error)
	Update(patient domain.Patient) (domain.Patient, error)
	Delete(id int) error
	CompleteEmptyAttributes(updatedPatient domain.Patient) (domain.Patient, error)
}
