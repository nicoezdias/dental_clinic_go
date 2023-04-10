package handler

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"dental_clinic_go/internal/domain"
	"dental_clinic_go/internal/patient"
	"dental_clinic_go/pkg/web"

	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	s patient.PatientService
}

// NewPatientHandler crea un nuevo controller de pacientes
func NewPatientHandler(s patient.PatientService) *patientHandler {
	return &patientHandler{s}
}

// GetByID godoc
// @Summary      Get a patient by Id
// @Description  Get a patient by Id from repository
// @Tags         patients
// @Produce      json
// @Param        token header string true "token"
// @Param        id   path      int  true  "Patient Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /patients/:id [get]
func (h *patientHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		patient, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		web.Success(c, 200, patient)
	}
}

// Post godoc
// @Summary      Create a new patient
// @Description  Create a new patient in repository
// @Tags         patients
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Patient true "Patient"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /patients [post]
func (h *patientHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var patient domain.Patient
		err := c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := h.validateEmptys(patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = h.validateAdmissionDate(patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(patient)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Put godoc
// @Summary      Update a patient by id
// @Description  Update a patient by id in repository
// @Tags         patients
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Patient true "Patient"
// @Param        id   path      int  true  "Patient Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /patients/:id [put]
func (h *patientHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		var patient domain.Patient
		err = c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		valid, err := h.validateEmptys(patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = h.validateAdmissionDate(patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, patient)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Patch godoc
// @Summary      Update a patient
// @Description  Update a patient by id in repository
// @Tags         patients
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Patient true "Patient"
// @Param        id   path      int  true  "Patient Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /patients/:id [patch]
func (h *patientHandler) Patch() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		var patient domain.Patient
		err = c.BindJSON(&patient)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		if patient.AdmissionDate != "" {
			valid, err := h.validateAdmissionDate(patient)
			if !valid {
				web.Failure(c, 400, err)
				return
			}
		}
		p, err := h.s.Update(id, patient)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Delete godoc
// @Summary      Delete a patient
// @Description  Delete a patient by id in repository
// @Tags         patients
// @Produce      json
// @Param        token header string true "token"
// @Param        id   path      int  true  "Patient Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /patients/:id [delete]
func (h *patientHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, fmt.Sprintf("patient %d deleted", id))
	}
}

/* ---------------------------------- Utils --------------------------------- */

// validateEmptys valida que los campos no esten vacios
func (h *patientHandler) validateEmptys(patient domain.Patient) (bool, error) {
	switch {
	case patient.Name == "":
		return false, errors.New("name can't be empty")
	case patient.LastName == "":
		return false, errors.New("last_name can't be empty")
	case patient.Dni == 0:
		return false, errors.New("dni can't be empty")
	case patient.Email == "":
		return false, errors.New("email can't be empty")
	case patient.AdmissionDate == "":
		return false, errors.New("admission_date can't be empty")
	}
	return true, nil
}

// validateExpiration valida que la fecha de expiracion sea valida
func (h *patientHandler) validateAdmissionDate(patient domain.Patient) (bool, error) {
	dates := strings.Split(patient.AdmissionDate, "-")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid admission_date, must be in format: dd-mm-yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid admission_date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[2] < 1 || list[2] > 31) && (list[1] < 1 || list[1] > 12) && (list[0] < 1 || list[0] > 9999)
	if condition {
		return false, errors.New("invalid admission_date, date must be between 1 and 31-12-9999")
	}
	return true, nil
}
