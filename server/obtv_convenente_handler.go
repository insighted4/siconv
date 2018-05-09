package server

import (
	"net/http"
	"path"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *server) CreateOBTVConvenenteHandler(c *gin.Context) {
	var obtvConvenente schema.OBTVConvenente
	if err := c.BindJSON(&obtvConvenente); err != nil {
		s.logger.Error(err)
		abort(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := s.service.CreateOBTVConvenente(&obtvConvenente)
	if err != nil {
		switch err {
		case storage.ErrAlreadyExists:
			abort(c, http.StatusUnprocessableEntity, err.Error())
		default:
			abort(c, http.StatusInternalServerError, err.Error())
		}

		s.logger.Error(err)
		return
	}

	location := path.Join(Prefix, "obtv-convenentes", id)
	c.Header("Location", location)
	c.Writer.WriteHeader(http.StatusCreated)
}

func (s *server) GetOBTVConvenenteHandler(c *gin.Context) {
	model, err := s.service.GetOBTVConvenente(c.Param("id"))
	switch err {
	case storage.ErrNotFound:
		abort(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
	case storage.ErrInvalidUUID:
		abort(c, http.StatusBadRequest, err.Error())
	case nil:
		location := path.Join(Prefix, "obtv-convenentes", model.ID)
		c.Header("Location", location)
		c.JSON(http.StatusOK, model)
	default:
		s.logger.Error(err)
		abort(c, http.StatusInternalServerError, err.Error())
	}
}

func (s *server) ListOBTVConvenenteHandler(c *gin.Context) {
	pagination := getPagination(c)
	models, total, err := s.service.ListOBTVConvenente(pagination)
	switch err {
	case nil:
		c.Header("X-Total-Count", strconv.Itoa(total))
		c.JSON(http.StatusOK, models)
	default:
		s.logger.Error(err)
		abort(c, http.StatusInternalServerError, err.Error())
	}
}
