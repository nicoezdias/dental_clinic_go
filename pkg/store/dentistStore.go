package store

import "dental_clinic_go/internal/domain"

type DentistStore interface {
	GetByID(id int) (domain.Dentist, error)
	GetByLicense(license string) (domain.Dentist, error)
	Create(dentist domain.Dentist) (domain.Dentist, error)
	Update(dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
	CompleteEmptyAttributes(updatedDentist domain.Dentist) (domain.Dentist, error)
}
