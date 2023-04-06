package dentist

import (
	"dental_clinic_go/internal/domain"
	"dental_clinic_go/pkg/store"
	"errors"
	"fmt"
)

type DentistRepository interface {
	GetByID(id int) (domain.Dentist, error)
	Create(p domain.Dentist) (domain.Dentist, error)
	Update(id int, updatedDentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type dentistRepository struct {
	storage store.DentistStore
}

// NewDentistRepository crea un nuevo repositorio
func NewDentistRepository(storage store.DentistStore) DentistRepository {
	return &dentistRepository{storage}
}

// GetByID busca un dentista por su id
func (r *dentistRepository) GetByID(id int) (domain.Dentist, error) {
	dentist, err := r.storage.GetByID(id)
	if err != nil {
		return domain.Dentist{}, errors.New(fmt.Sprintf("dentist %d not found", id))
	}
	return dentist, nil
}

// Create agrega un nuevo dentista
func (r *dentistRepository) Create(d domain.Dentist) (domain.Dentist, error) {
	_, err := r.storage.GetByLicense(d.License)
	if err == nil {
		return domain.Dentist{}, errors.New("license already exists")
	}
	dentist, err := r.storage.Create(d)
	if err != nil {
		return domain.Dentist{}, errors.New("error creating dentist")
	}
	return dentist, nil
}

// Update actualiza un dentista
func (r *dentistRepository) Update(id int, updatedDentist domain.Dentist) (domain.Dentist, error) {
	dentist, err := r.storage.GetByLicense(updatedDentist.License)
	if err == nil && dentist.Id != id {
		return domain.Dentist{}, errors.New("license already exists")
	}
	updatedDentist.Id = id
	p, err := r.storage.Update(updatedDentist)
	if err != nil {
		return domain.Dentist{}, errors.New("error updating dentist")
	}
	return p, nil
}

// Delete busca un dentista por su id y lo elimina
func (r *dentistRepository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
