package patient

import (
	"dental_clinic_go/internal/domain"
)

type PatientService interface {
	GetByID(id int) (domain.Patient, error)
	Create(p domain.Patient) (domain.Patient, error)
	Update(id int, updatedPatient domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type patientService struct {
	r PatientRepository
}

// NewService crea un nuevo servicio
func NewPatientService(r PatientRepository) PatientService {
	return &patientService{r}
}

// GetByID busca un paciente por su id
func (s *patientService) GetByID(id int) (domain.Patient, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

// Create agrega un nuevo paciente
func (s *patientService) Create(p domain.Patient) (domain.Patient, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

// UpdatePatient actualiza un paciente
func (s *patientService) Update(id int, updatedPatient domain.Patient) (domain.Patient, error) {
	p, err := s.r.Update(id, updatedPatient)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

// Delete busca un paciente por su id y lo elimina
func (s *patientService) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
