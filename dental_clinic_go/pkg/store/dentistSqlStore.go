package store

import (
	"database/sql"
	"dental_clinic_go/internal/domain"
)

type dentistSqlStore struct {
	DB *sql.DB
}

// NewSqlStore crea un nuevo store de dentists
func NewDentistSqlStore(db *sql.DB) DentistStore {
	return &dentistSqlStore{db}
}

// GetByID devuelve un dentista por su id
func (s *dentistSqlStore) GetByID(id int) (domain.Dentist, error) {
	var dentistReturn domain.Dentist

	query := "SELECT * FROM dentist WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&dentistReturn.Id, &dentistReturn.Name, &dentistReturn.LastName, &dentistReturn.License)

	if err != nil {
		return domain.Dentist{}, err
	}
	return dentistReturn, nil
}

// GetByLicense devuelve un dentista por su matricula
func (s *dentistSqlStore) GetByLicense(license string) (domain.Dentist, error) {
	var dentistReturn domain.Dentist

	query := "SELECT * FROM dentist WHERE license = ?;"
	row := s.DB.QueryRow(query, license)
	err := row.Scan(&dentistReturn.Id, &dentistReturn.Name, &dentistReturn.LastName, &dentistReturn.License)

	if err != nil {
		return domain.Dentist{}, err
	}
	return dentistReturn, nil
}

// Create agrega un nuevo dentista
func (s *dentistSqlStore) Create(dentist domain.Dentist) (domain.Dentist, error) {
	stmt, err := s.DB.Prepare("INSERT INTO dentist(name, last_name, license) VALUES( ?, ?, ?)")
	if err != nil {
		return domain.Dentist{}, err
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(dentist.Name, dentist.LastName, dentist.License)
	if err != nil {
		return domain.Dentist{}, err
	}
	insertedId, _ := result.LastInsertId()
	dentist.Id = int(insertedId)
	return dentist, nil
}

// Update actualiza un dentista
func (s *dentistSqlStore) Update(dentist domain.Dentist) (domain.Dentist, error) {
	dentistUpdated, err := s.CompleteEmptyAttributes(dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	stmt, err := s.DB.Prepare("UPDATE dentist SET name = ?, last_name = ?, license = ? WHERE id = ?;")
	if err != nil {
		return domain.Dentist{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(dentistUpdated.Name, dentistUpdated.LastName, dentistUpdated.License, dentist.Id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentistUpdated, nil
}

// Delete elimina un dentista
func (s *dentistSqlStore) Delete(id int) error {
	stmt := "DELETE FROM dentist WHERE id = ?"
	_, err := s.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

// completeEmptyAttributes compara dos dentistas y se queda con los campos diferentes
func (s *dentistSqlStore) CompleteEmptyAttributes(updatedDentist domain.Dentist) (domain.Dentist, error) {
	d, err := s.GetByID(updatedDentist.Id)
	if err != nil {
		return updatedDentist, err
	}
	if updatedDentist.Name != "" {
		d.Name = updatedDentist.Name
	}
	if updatedDentist.LastName != "" {
		d.LastName = updatedDentist.LastName
	}
	if updatedDentist.License != "" {
		d.License = updatedDentist.License
	}
	return d, nil
}
