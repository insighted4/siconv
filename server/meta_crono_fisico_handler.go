package server

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
)

func (s *server) CreateMetaCronoFisicoHandler(c *gin.Context) {
	var metaCronoFisico schema.MetaCronoFisico
	if err := c.BindJSON(&metaCronoFisico); err != nil {
		s.logger.Error(err)
		abort(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := s.service.CreateMetaCronoFisico(&metaCronoFisico)
	if err != nil {
		switch err {
		case siconv.ErrAlreadyExists:
			abort(c, http.StatusUnprocessableEntity, err.Error())
		default:
			abort(c, http.StatusInternalServerError, err.Error())
		}

		s.logger.Error(err)
		return
	}

	location := path.Join(Prefix, "meta-crono-fisicos", id)
	c.Header("Location", location)
	c.Writer.WriteHeader(http.StatusCreated)
}

func (s *server) GetMetaCronoFisicoHandler(c *gin.Context) {
	model, err := s.service.GetMetaCronoFisico(c.Param("id"))
	switch err {
	case siconv.ErrNotFound:
		abort(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
	case siconv.ErrInvalidUUID:
		abort(c, http.StatusBadRequest, err.Error())
	case nil:
		location := path.Join(Prefix, "meta-crono-fisicos", model.ID)
		c.Header("Location", location)
		c.JSON(http.StatusOK, model)
	default:
		s.logger.Error(err)
		abort(c, http.StatusInternalServerError, err.Error())
	}
}
