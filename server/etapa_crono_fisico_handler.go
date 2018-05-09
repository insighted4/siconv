package server

import (
	"net/http"
	"path"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *server) CreateEtapaCronoFisicoHandler(c *gin.Context) {
	var model schema.EtapaCronoFisico
	if err := c.BindJSON(&model); err != nil {
		s.logger.Error(err)
		abort(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.service.Create(&model); err != nil {
		switch err {
		case storage.ErrAlreadyExists:
			abort(c, http.StatusUnprocessableEntity, err.Error())
		default:
			abort(c, http.StatusInternalServerError, err.Error())
		}

		s.logger.Error(err)
		return
	}

	location := path.Join(Prefix, "etapa-crono-fisicos", strconv.Itoa(model.GetID()))
	c.Header("Location", location)
	c.Writer.WriteHeader(http.StatusCreated)
}

func (s *server) GetEtapaCronoFisicoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		abort(c, http.StatusBadRequest, "ID should be an integer")
		return
	}

	model := &schema.EtapaCronoFisico{StorageModel: schema.StorageModel{ID: id}}
	switch err := s.service.Get(model); err {
	case storage.ErrNotFound:
		abort(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
	case storage.ErrInvalidID:
		abort(c, http.StatusBadRequest, err.Error())
	case nil:
		location := path.Join(Prefix, "etapa-crono-fisicos", strconv.Itoa(model.GetID()))
		c.Header("Location", location)
		c.JSON(http.StatusOK, model)
	default:
		s.logger.Error(err)
		abort(c, http.StatusInternalServerError, err.Error())
	}
}

func (s *server) ListEtapaCronoFisicoHandler(c *gin.Context) {
	pagination := getPagination(c)

	models := []*schema.EtapaCronoFisico{nil}
	total, err := s.service.List(&models, pagination)
	switch err {
	case nil:
		c.Header("X-Total-Count", strconv.Itoa(total))
		c.JSON(http.StatusOK, models)
	default:
		s.logger.Error(err)
		abort(c, http.StatusInternalServerError, err.Error())
	}
}
