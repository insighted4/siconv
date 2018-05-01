package server

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
)

func (s *server) CreateHistoricoSituacaoHandler(c *gin.Context) {
	var historicoSituacao schema.HistoricoSituacao
	if err := c.BindJSON(&historicoSituacao); err != nil {
		s.logger.Error(err)
		abort(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := s.service.CreateHistoricoSituacao(&historicoSituacao)
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

	location := path.Join(Prefix, "historico-situacoes", id)
	c.Header("Location", location)
	c.Writer.WriteHeader(http.StatusCreated)
}

func (s *server) GetHistoricoSituacaoHandler(c *gin.Context) {
	model, err := s.service.GetHistoricoSituacao(c.Param("id"))
	switch err {
	case siconv.ErrNotFound:
		abort(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
	case siconv.ErrInvalidUUID:
		abort(c, http.StatusBadRequest, err.Error())
	case nil:
		location := path.Join(Prefix, "historico-situacoes", model.ID)
		c.Header("Location", location)
		c.JSON(http.StatusOK, model)
	default:
		s.logger.Error(err)
		abort(c, http.StatusInternalServerError, err.Error())
	}
}