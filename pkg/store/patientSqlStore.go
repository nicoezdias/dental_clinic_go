package store

import (
	"database/sql"
	"dental_clinic_go/internal/domain"

	"time"
)

type patientSqlStore struct {
	DB *sql.DB
}

// NewSqlStore crea un nuevo store de patients
func NewPatientSqlStore(db *sql.DB) PatientStore {
	return &patientSqlStore{db}
}

// GetByID devuelve un paciente por su id
func (s *patientSqlStore) GetByID(id int) (domain.Patient, error) {
	var patientReturn domain.Patient
	query := "SELECT * FROM patient WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&patientReturn.Id, &patientReturn.Name, &patientReturn.LastName, &patientReturn.Domicilio, &patientReturn.Dni, &patientReturn.Email, &patientReturn.AdmissionDate)
	if err != nil {
		return domain.Patient{}, err
	}
	return patientReturn, nil
}

// GetByDni devuelve un paciente por su dni
func (s *patientSqlStore) GetByDni(dni int) (domain.Patient, error) {
	var patientReturn domain.Patient
	query := "SELECT * FROM patient WHERE dni = ?;"
	row := s.DB.QueryRow(query, dni)
	err := row.Scan(&patientReturn.Id, &patientReturn.Name, &patientReturn.LastName, &patientReturn.Domicilio, &patientReturn.Dni, &patientReturn.Email, &patientReturn.AdmissionDate)
	if err != nil {
		return domain.Patient{}, err
	}
	return patientReturn, nil
}

// Create agrega un nuevo paciente
func (s *patientSqlStore) Create(patient domain.Patient) (domain.Patient, error) {
	stmt, err := s.DB.Prepare("INSERT INTO patient (name, last_name, domicilio, dni, email, admission_date) VALUES (?, ?, ?, ?, ?, ?);")
	if err != nil {
		return domain.Patient{}, err
	}
	defer stmt.Close()
	date, err := time.Parse("2006-01-02", patient.AdmissionDate)
	if err != nil {
		return domain.Patient{}, err
	}
	result, err := stmt.Exec(patient.Name, patient.LastName, patient.Domicilio, patient.Dni, patient.Email, date)
	if err != nil {
		return domain.Patient{}, err
	}
	insertedId, _ := result.LastInsertId()
	patient.Id = int(insertedId)
	return patient, nil
}

// Update actualiza un paciente
func (s *patientSqlStore) Update(patient domain.Patient) (domain.Patient, error) {
	patientUpdated, err := s.CompleteEmptyAttributes(patient)
	stmt, err := s.DB.Prepare("UPDATE patient SET name = ?, last_name = ?, domicilio = ?, dni = ?, email = ?, admission_date = ? WHERE id = ?;")
	if err != nil {
		return domain.Patient{}, err
	}
	defer stmt.Close()
	date, err := time.Parse("2006-01-02", patient.AdmissionDate)
	if err != nil {
		return domain.Patient{}, err
	}
	_, err = stmt.Exec(patientUpdated.Name, patientUpdated.LastName, patientUpdated.Domicilio, patientUpdated.Dni, patientUpdated.Email, date, patientUpdated.Id)
	if err != nil {
		return domain.Patient{}, err
	}
	return patientUpdated, nil
}

// Delete elimina un paciente
func (s *patientSqlStore) Delete(id int) error {
	stmt := "DELETE FROM patient WHERE id = ?"
	_, err := s.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

// completeEmptyAttributes compara dos pacientes y se queda con los campos diferentes
func (s *patientSqlStore) CompleteEmptyAttributes(updatedPatient domain.Patient) (domain.Patient, error) {
	p, err := s.GetByID(updatedPatient.Id)
	if err != nil {
		return updatedPatient, err
	}
	if updatedPatient.Name != "" {
		p.Name = updatedPatient.Name
	}
	if updatedPatient.LastName != "" {
		p.LastName = updatedPatient.LastName
	}
	if updatedPatient.Domicilio != "" {
		p.Domicilio = updatedPatient.Domicilio
	}
	if updatedPatient.Dni != 0 {
		p.Dni = updatedPatient.Dni
	}
	if updatedPatient.Email != "" {
		p.Email = updatedPatient.Email
	}
	if updatedPatient.AdmissionDate != "" {
		p.AdmissionDate = updatedPatient.AdmissionDate
	}
	return p, nil
}
