package store

import (
	"database/sql"
	"dental_clinic_go/internal/domain"

	"time"
)

type appointmentSqlStore struct {
	DB *sql.DB
}

// NewSqlStore crea un nuevo store de appointments
func NewAppointmentSqlStore(db *sql.DB) AppointmentStore {
	return &appointmentSqlStore{db}
}

// GetByID devuelve un turno por su id
func (s *appointmentSqlStore) GetByID(id int) (domain.Appointment, error) {
	var appointmentReturn domain.Appointment
	query := "SELECT appointment.id, appointment.date, appointment.hour, patient.*, dentist.* FROM appointment INNER JOIN patient ON appointment.patient_id = patient.id INNER JOIN dentist ON appointment.dentist_id = dentist.id WHERE appointment.id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&appointmentReturn.Id, &appointmentReturn.Date, &appointmentReturn.Hour, &appointmentReturn.Patient.Id, &appointmentReturn.Patient.Name, &appointmentReturn.Patient.LastName, &appointmentReturn.Patient.Domicilio, &appointmentReturn.Patient.Dni, &appointmentReturn.Patient.Email, &appointmentReturn.Patient.AdmissionDate, &appointmentReturn.Dentist.Id, &appointmentReturn.Dentist.Name, &appointmentReturn.Dentist.LastName, &appointmentReturn.Dentist.License)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointmentReturn, nil
}

// GetByDni devuelve los turnos filtrando por un dni del paciente
func (s *appointmentSqlStore) GetByDni(dni int) ([]domain.Appointment, error) {
	var appointments []domain.Appointment

	query := "SELECT appointment.id, appointment.date, appointment.hour, patient.*, dentist.* FROM appointment INNER JOIN patient ON appointment.patient_id = patient.id INNER JOIN dentist ON appointment.dentist_id = dentist.id WHERE patient.dni = ?"
	rows, err := s.DB.Query(query, dni)
	if err != nil {
		return []domain.Appointment{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var appointmentReturn domain.Appointment
		err := rows.Scan(&appointmentReturn.Id, &appointmentReturn.Date, &appointmentReturn.Hour, &appointmentReturn.Patient.Id, &appointmentReturn.Patient.Name, &appointmentReturn.Patient.LastName, &appointmentReturn.Patient.Domicilio, &appointmentReturn.Patient.Dni, &appointmentReturn.Patient.Email, &appointmentReturn.Patient.AdmissionDate, &appointmentReturn.Dentist.Id, &appointmentReturn.Dentist.Name, &appointmentReturn.Dentist.LastName, &appointmentReturn.Dentist.License)
		if err != nil {
			return []domain.Appointment{}, err
		}
		appointments = append(appointments, appointmentReturn)
	}
	if err = rows.Err(); err != nil {
		return []domain.Appointment{}, err
	}
	return appointments, nil
}

// Create agrega un nuevo turno
func (s *appointmentSqlStore) Create(appointment domain.Appointment) (domain.Appointment, error) {
	stmt, err := s.DB.Prepare("INSERT INTO appointment (date, hour, patient_id, dentist_id) VALUES (?, ?, ?, ?);")
	if err != nil {
		return domain.Appointment{}, err
	}
	defer stmt.Close()
	date, err := time.Parse("2006-01-02", appointment.Date)
	if err != nil {
		return domain.Appointment{}, err
	}
	t, err := time.Parse("15:04:05", appointment.Hour)
	if err != nil {
		return domain.Appointment{}, err
	}
	hour := time.Date(1970, 1, 1, t.Hour(), t.Minute(), t.Second(), 0, time.UTC).Format("15:04:05")
	result, err := stmt.Exec(date, hour, appointment.Patient.Id, appointment.Dentist.Id)
	if err != nil {
		return domain.Appointment{}, err
	}
	insertedId, _ := result.LastInsertId()
	appointment.Id = int(insertedId)
	return appointment, nil
}

// Update actualiza un turno
func (s *appointmentSqlStore) Update(appointment domain.Appointment) (bool, bool, domain.Appointment, error) {
	patientFlag, dentistFlag, appointmentUpdated, err := s.CompleteEmptyAttributes(appointment)
	if err != nil {
		return false, false, domain.Appointment{}, err
	}
	stmt, err := s.DB.Prepare("UPDATE appointment SET date = ?, hour = ?, patient_id = ?, dentist_id = ? WHERE id = ?;")
	if err != nil {
		return false, false, domain.Appointment{}, err
	}
	defer stmt.Close()
	date, err := time.Parse("2006-01-02", appointment.Date)
	if err != nil {
		return false, false, domain.Appointment{}, err
	}
	t, err := time.Parse("15:04:05", appointment.Hour)
	if err != nil {
		return false, false, domain.Appointment{}, err
	}
	hour := time.Date(1970, 1, 1, t.Hour(), t.Minute(), t.Second(), 0, time.UTC).Format("15:04:05")
	_, err = stmt.Exec(date, hour, appointmentUpdated.Patient.Id, appointmentUpdated.Dentist.Id, appointmentUpdated.Id)
	if err != nil {
		return false, false, domain.Appointment{}, err
	}
	return patientFlag, dentistFlag, appointmentUpdated, nil
}

// Delete elimina un turno
func (s *appointmentSqlStore) Delete(id int) error {
	stmt := "DELETE FROM appointment WHERE id = ?"
	_, err := s.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

// completeEmptyAttributes compara dos turnos y se queda con los campos diferentes
func (s *appointmentSqlStore) CompleteEmptyAttributes(updatedAppointment domain.Appointment) (bool, bool, domain.Appointment, error) {
	a, err := s.GetByID(updatedAppointment.Id)
	if err != nil {
		return false, false, domain.Appointment{}, err
	}
	patientFlag := false
	dentistFlag := false
	if updatedAppointment.Date != "" {
		a.Date = updatedAppointment.Date
	}
	if updatedAppointment.Hour != "" {
		a.Hour = updatedAppointment.Hour
	}
	if (updatedAppointment.Patient != domain.Patient{} && updatedAppointment.Patient.Id != 0) {
		if a.Patient.Id != updatedAppointment.Patient.Id {
			patientFlag = true
		}
		a.Patient = updatedAppointment.Patient
	}
	if (updatedAppointment.Dentist != domain.Dentist{} && updatedAppointment.Dentist.Id != 0) {
		if a.Dentist.Id != updatedAppointment.Dentist.Id {
			dentistFlag = true
		}
		a.Dentist = updatedAppointment.Dentist
	}
	return patientFlag, dentistFlag, a, nil
}
