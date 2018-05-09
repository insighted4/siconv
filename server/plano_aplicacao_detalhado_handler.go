package server

import (
	"net/http"
	"path"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *server) CreatePlanoAplicacaoDetalhadoHandler(c *gin.Context) {
	var planoAplicacaoDetalhado schema.PlanoAplicacaoDetalhado
	if err := c.BindJSON(&planoAplicacaoDetalhado); err != nil {
		s.logger.Error(err)
		abort(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := s.service.CreatePlanoAplicacaoDetalhado(&planoAplicacaoDetalhado)
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

	location := path.Join(Prefix, "plano-aplicacao-detalhados", id)
	c.Header("Location", location)
	c.Writer.WriteHeader(http.StatusCreated)
}

func (s *server) GetPlanoAplicacaoDetalhadoHandler(c *gin.Context) {
	model, err := s.service.GetPlanoAplicacaoDetalhado(c.Param("id"))
	switch err {
	case storage.ErrNotFound:
		abort(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
	case storage.ErrInvalidUUID:
		abort(c, http.StatusBadRequest, err.Error())
	case nil:
		location := path.Join(Prefix, "plano-aplicacao-detalhados", model.ID)
		c.Header("Location", location)
		c.JSON(http.StatusOK, model)
	default:
		s.logger.Error(err)
		abort(c, http.StatusInternalServerError, err.Error())
	}
}

func (s *server) ListPlanoAplicacaoDetalhadoHandler(c *gin.Context) {
	pagination := getPagination(c)
	models, total, err := s.service.ListPlanoAplicacaoDetalhado(pagination)
	switch err {
	case nil:
		c.Header("X-Total-Count", strconv.Itoa(total))
		c.JSON(http.StatusOK, models)
	default:
		s.logger.Error(err)
		abort(c, http.StatusInternalServerError, err.Error())
	}
}
