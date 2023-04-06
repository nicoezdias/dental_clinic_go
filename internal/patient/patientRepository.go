package patient

import (
	"dental_clinic_go/internal/domain"
	"dental_clinic_go/pkg/store"
	"errors"
	"fmt"
)

type PatientRepository interface {
	GetByID(id int) (domain.Patient, error)
	Create(p domain.Patient) (domain.Patient, error)
	Update(id int, updatedPatient domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type patientRepository struct {
	storage store.PatientStore
}

// NewPatientRepository crea un nuevo repositorio
func NewPatientRepository(storage store.PatientStore) PatientRepository {
	return &patientRepository{storage}
}

// GetByID busca un paciente por su id
func (r *patientRepository) GetByID(id int) (domain.Patient, error) {
	patient, err := r.storage.GetByID(id)
	if err != nil {
		return domain.Patient{}, errors.New(fmt.Sprintf("patient %d not found", id))
	}
	return patient, nil
}

// Create agrega un nuevo paciente
func (r *patientRepository) Create(p domain.Patient) (domain.Patient, error) {
	_, err := r.storage.GetByDni(p.Dni)
	if err == nil {
		return domain.Patient{}, errors.New("dni already exists")
	}
	patient, err := r.storage.Create(p)
	if err != nil {
		return domain.Patient{}, errors.New("error creating patient")
	}
	return patient, nil
}

// Update actualiza un paciente
func (r *patientRepository) Update(id int, updatedPatient domain.Patient) (domain.Patient, error) {
	patient, err := r.storage.GetByDni(updatedPatient.Dni)
	if err == nil && patient.Id != id {
		return domain.Patient{}, errors.New("license already exists")
	}
	updatedPatient.Id = id
	p, err := r.storage.Update(updatedPatient)
	if err != nil {
		return domain.Patient{}, errors.New("error updating patient")
	}
	return p, nil
}

// Delete busca un paciente por su id y lo elimina
func (r *patientRepository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
