package appointment

import (
	"dental_clinic_go/internal/domain"
)

type AppointmentService interface {
	GetByID(id int) (domain.Appointment, error)
	GetByDni(id int) ([]domain.Appointment, error)
	Create(a domain.Appointment) (domain.Appointment, error)
	CreateByDniAndLicense(dni int, license string, appointment domain.Appointment) (domain.Appointment, error)
	Update(id int, updatedAppointment domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}

type appointmentService struct {
	r AppointmentRepository
}

// NewService crea un nuevo servicio
func NewAppointmentService(r AppointmentRepository) AppointmentService {
	return &appointmentService{r}
}

// GetByID busca un turno por su id
func (s *appointmentService) GetByID(id int) (domain.Appointment, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	return p, nil
}

// GetByID busca un turno por su id
func (s *appointmentService) GetByDni(dni int) ([]domain.Appointment, error) {
	p, err := s.r.GetByDni(dni)
	if err != nil {
		return []domain.Appointment{}, err
	}
	return p, nil
}

// Create agrega un nuevo turno
func (s *appointmentService) Create(a domain.Appointment) (domain.Appointment, error) {
	p, err := s.r.Create(a)
	if err != nil {
		return domain.Appointment{}, err
	}
	return p, nil
}

// CreateByDniAndLicense agrega un nuevo turno por medio de el dni del paciente y la matricula del dentista
func (s *appointmentService) CreateByDniAndLicense(dni int, license string, appointment domain.Appointment) (domain.Appointment, error) {
	p, err := s.r.CreateByDniAndLicense(dni, license, appointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	return p, nil
}

// UpdateAppointment actualiza un turno
func (s *appointmentService) Update(id int, updatedAppointment domain.Appointment) (domain.Appointment, error) {
	p, err := s.r.Update(id, updatedAppointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	return p, nil
}

// Delete busca un turno por su id y lo elimina
func (s *appointmentService) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
