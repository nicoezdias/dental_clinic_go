package handler

import (
	"errors"
	"fmt"
	"strconv"

	"dental_clinic_go/internal/dentist"
	"dental_clinic_go/internal/domain"
	"dental_clinic_go/pkg/web"

	"github.com/gin-gonic/gin"
)

type dentistHandler struct {
	s dentist.Service
}

// NewDentistHandler crea un nuevo controller de dentistas
func NewDentistHandler(s dentist.Service) *dentistHandler {
	return &dentistHandler{s}
}

// GetByID godoc
// @Summary      Get a dentist by Id
// @Description  Get a dentist by Id from repository
// @Tags         dentists
// @Produce      json
// @Param        token header string true "token"
// @Param        id   path      int  true  "Dentist Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /dentists/:id [get]
func (h *dentistHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		dentist, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		web.Success(c, 200, dentist)
	}
}

// Post godoc
// @Summary      Create a new dentist
// @Description  Create a new dentist in repository
// @Tags         dentists
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Dentist true "Dentist"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /dentists [post]
func (h *dentistHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentist domain.Dentist
		err := c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := h.validateEmptys(dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(dentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Put godoc
// @Summary      Update a dentist by id
// @Description  Update a dentist by id in repository
// @Tags         dentists
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Dentist true "Dentist"
// @Param        id   path      int  true  "Dentist Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /dentists/:id [put]
func (h *dentistHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		var dentist domain.Dentist
		err = c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		valid, err := h.validateEmptys(dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, dentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Patch godoc
// @Summary      Update a dentist
// @Description  Update a dentist by id in repository
// @Tags         dentists
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Dentist true "Dentist"
// @Param        id   path      int  true  "Dentist Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /dentists/:id [patch]
func (h *dentistHandler) Patch() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		var dentist domain.Dentist
		err = c.BindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, dentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Delete godoc
// @Summary      Delete a dentist
// @Description  Delete a dentist by id in repository
// @Tags         dentists
// @Produce      json
// @Param        token header string true "token"
// @Param        id   path      int  true  "Dentist Id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /dentists/:id [delete]
func (h *dentistHandler) Delete() gin.HandlerFunc {
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
		web.Success(c, 204, fmt.Sprintf("dentist %d deleted", id))
	}
}

/* ---------------------------------- Utils --------------------------------- */

// validateEmptys valida que los campos no esten vacios
func (h *dentistHandler) validateEmptys(dentist domain.Dentist) (bool, error) {
	switch {
	case dentist.Name == "":
		return false, errors.New("name can't be empty")
	case dentist.LastName == "":
		return false, errors.New("last_name can't be empty")
	case dentist.License == "":
		return false, errors.New("license can't be empty")
	}
	return true, nil
}
