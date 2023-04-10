package dentist

import (
	"dental_clinic_go/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Dentist, error)
	Create(p domain.Dentist) (domain.Dentist, error)
	Update(id int, updatedDentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type service struct {
	r DentistRepository
}

// NewService crea un nuevo servicio
func NewDentistService(r DentistRepository) Service {
	return &service{r}
}

// GetByID busca un dentista por su id
func (s *service) GetByID(id int) (domain.Dentist, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return p, nil
}

// Create agrega un nuevo dentista
func (s *service) Create(p domain.Dentist) (domain.Dentist, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Dentist{}, err
	}
	return p, nil
}

// UpdateDentist actualiza un dentista
func (s *service) Update(id int, updatedDentist domain.Dentist) (domain.Dentist, error) {
	p, err := s.r.Update(id, updatedDentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return p, nil
}

// Delete busca un dentista por su id y lo elimina
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
