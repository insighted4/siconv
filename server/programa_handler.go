package server

import (
	"net/http"
	"path"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
)

func (s *server) CreateProgramaHandler(c *gin.Context) {
	var programa schema.Programa
	if err := c.BindJSON(&programa); err != nil {
		s.logger.Error(err)
		abort(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := s.service.CreatePrograma(&programa)
	switch err {
	case siconv.ErrAlreadyExists:
		abort(c, http.StatusUnprocessableEntity, err.Error())
	case nil:
		location := path.Join(Prefix, "programas", id)
		c.Header("Location", location)
		c.Writer.WriteHeader(http.StatusCreated)
	default:
		s.logger.Error(err)
		abort(c, http.StatusInternalServerError, err.Error())
	}
}

func (s *server) GetProgramaHandler(c *gin.Context) {
	programa, err := s.service.GetPrograma(c.Param("id"))
	switch err {
	case siconv.ErrNotFound:
		abort(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
	case siconv.ErrInvalidUUID:
		abort(c, http.StatusBadRequest, err.Error())
	case nil:
		location := path.Join(Prefix, "programas", programa.ID)
		c.Header("Location", location)
		c.JSON(http.StatusOK, programa)
	default:
		s.logger.Error(err)
		abort(c, http.StatusInternalServerError, err.Error())
	}
}

func (s *server) ListProgramaHandler(c *gin.Context) {
	pagination := getPagination(c)
	idPrograma := c.DefaultQuery("id_programa", "")

	models, total, err := s.service.ListPrograma(idPrograma, pagination)
	switch err {
	case nil:
		c.Header("X-Total-Count", strconv.Itoa(total))
		c.JSON(http.StatusOK, models)
	default:
		s.logger.Error(err)
		abort(c, http.StatusInternalServerError, err.Error())
	}
}
