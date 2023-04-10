package handler

import (
	"errors"
	"fmt"

	"strconv"
	"strings"

	"dental_clinic_go/internal/appointment"
	"dental_clinic_go/internal/domain"
	"dental_clinic_go/pkg/web"

	"github.com/gin-gonic/gin"
)

type appointmentHandler struct {
	s appointment.AppointmentService
}

// NewAppointmentHandler crea un nuevo controller de pacientes
func NewAppointmentHandler(s appointment.AppointmentService) *appointmentHandler {
	return &appointmentHandler{s}
}

// GetByID godoc
// @Summary      Get a appointment by Id
// @Description  Get a appointment by Id from repository
// @Tags         appointments
// @Produce      json
// @Param        token header string true "token"
// @Param        id   path      int  true  "Appointment Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /appointments/:id [get]
func (h *appointmentHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		appointment, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, appointment)
	}
}

// GetByDni godoc
// @Summary      Get a appointments by patient.dni
// @Description  Get a appointments by patient.dni from repository
// @Tags         appointments
// @Produce      json
// @Param        token header string true "token"
// @Param        dni   path      int  true  "Patient Dni"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /appointments/dni/:dni [get]
func (h *appointmentHandler) GetByDni() gin.HandlerFunc {
	return func(c *gin.Context) {
		dniParam := c.Param("dni")
		dni, err := strconv.Atoi(dniParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid dni"))
			return
		}
		appointment, err := h.s.GetByDni(dni)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, appointment)
	}
}

// Post godoc
// @Summary      Create a new appointment
// @Description  Create a new appointment in repository
// @Tags         appointments
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Appointment true "Appointment"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /appointments [post]
func (h *appointmentHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var appointment domain.Appointment
		err := c.ShouldBindJSON(&appointment)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := h.validateEmptys(appointment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = h.validateDate(appointment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = h.validateHour(appointment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(appointment)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// PostByDniAndLicense godoc
// @Summary      Create a new appointment through the patient's ID and the dentist's license
// @Description  Create a new appointment through the patient's ID and the dentist's license in repository
// @Tags         appointments
// @Produce      json
// @Param        token header string true "token"
// @Param        dni   query      int  true  "DNI"
// @Param        license   query      string  true  "License"
// @Param        body body domain.Appointment true "Appointment"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /appointments/dni/license [post]
func (h *appointmentHandler) PostByDniAndLicense() gin.HandlerFunc {
	return func(c *gin.Context) {
		var appointment domain.Appointment
		err := c.ShouldBindJSON(&appointment)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		dniParam := c.Query("dni")
		dni, err := strconv.Atoi(dniParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid price"))
			return
		}
		license := c.Query("license")
		valid, err := h.validateDate(appointment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = h.validateHour(appointment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.CreateByDniAndLicense(dni, license, appointment)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Put godoc
// @Summary      Update a appointment by id
// @Description  Update a appointment by id in repository
// @Tags         appointments
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Appointment true "Appointment"
// @Param        id   path      int  true  "Appointment Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /appointments/:id [put]
func (h *appointmentHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		var appointment domain.Appointment
		err = c.ShouldBindJSON(&appointment)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		valid, err := h.validateEmptys(appointment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = h.validateDate(appointment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = h.validateHour(appointment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, appointment)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Patch godoc
// @Summary      Update a appointment
// @Description  Update a appointment by id in repository
// @Tags         appointments
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Appointment true "Appointment"
// @Param        id   path      int  true  "Appointment Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /appointments/:id [patch]
func (h *appointmentHandler) Patch() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		var appointment domain.Appointment
		err = c.BindJSON(&appointment)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		if appointment.Date != "" {
			valid, err := h.validateDate(appointment)
			if !valid {
				web.Failure(c, 400, err)
				return
			}
		}
		if appointment.Hour != "" {
			valid, err := h.validateHour(appointment)
			if !valid {
				web.Failure(c, 400, err)
				return
			}
		}
		p, err := h.s.Update(id, appointment)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Delete godoc
// @Summary      Delete a appointment
// @Description  Delete a appointment by id in repository
// @Tags         appointments
// @Produce      json
// @Param        token header string true "token"
// @Param        id   path      int  true  "Appointment Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /appointments/:id [delete]
func (h *appointmentHandler) Delete() gin.HandlerFunc {
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
		web.Success(c, 204, fmt.Sprintf("user %d deleted", id))
	}
}

/* ---------------------------------- Utils --------------------------------- */

// validateEmptys valida que los campos no esten vacios
func (h *appointmentHandler) validateEmptys(appointment domain.Appointment) (bool, error) {
	switch {
	case appointment.Description == "":
		return false, errors.New("Description can't be empty")
	case appointment.Patient == domain.Patient{}:
		return false, errors.New("Patient can't be empty")
	case appointment.Patient.Id == 0:
		return false, errors.New("Patient.id can't be empty")
	case appointment.Dentist == domain.Dentist{}:
		return false, errors.New("Dentist can't be empty")
	case appointment.Dentist.Id == 0:
		return false, errors.New("Dentist.id can't be empty")
	}
	return true, nil
}

// validateExpiration valida que la fecha de expiracion sea valida
func (h *appointmentHandler) validateDate(appointment domain.Appointment) (bool, error) {
	if appointment.Date == "" {
		return false, errors.New("date can't be empty")
	}
	dates := strings.Split(appointment.Date, "-")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid date, must be in format: dd-mm-yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[2] < 1 || list[2] > 31) && (list[1] < 1 || list[1] > 12) && (list[0] < 1 || list[0] > 9999)
	if condition {
		return false, errors.New("invalid admission_date, date must be between 1 and 31-12-9999")
	}
	return true, nil
}

// validateExpiration valida que la fecha de expiracion sea valida
func (h *appointmentHandler) validateHour(appointment domain.Appointment) (bool, error) {
	if appointment.Hour == "" {
		return false, errors.New("hour can't be empty")
	}
	hours := strings.Split(appointment.Hour, ":")
	list := []int{}
	if len(hours) != 3 {
		return false, errors.New("invalid hour, must be in format: hh-mm-ss")
	}
	for value := range hours {
		number, err := strconv.Atoi(hours[value])
		if err != nil {
			return false, errors.New("invalid hour, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[2] < 0 || list[2] > 59) && (list[1] < 0 || list[1] > 59) && (list[0] < 0 || list[0] > 23)
	if condition {
		return false, errors.New("invalid hour, date must be between 00:00:00 and 23:59:59")
	}
	return true, nil
}
