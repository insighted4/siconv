package server

import (
	"net/http"
	"path"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *server) CreateProgramaPropostaHandler(c *gin.Context) {
	var programaProposta schema.ProgramaProposta
	if err := c.BindJSON(&programaProposta); err != nil {
		s.logger.Error(err)
		abort(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := s.service.CreateProgramaProposta(&programaProposta)
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

	location := path.Join(Prefix, "programa-propostas", id)
	c.Header("Location", location)
	c.Writer.WriteHeader(http.StatusCreated)
}

func (s *server) ListProgramaPropostaHandler(c *gin.Context) {
	pagination := getPagination(c)

	programa, err := s.service.GetPrograma(c.Param("id"))
	if err != nil {
		switch err {
		case storage.ErrNotFound:
			abort(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		case storage.ErrInvalidUUID:
			abort(c, http.StatusBadRequest, err.Error())
		default:
			s.logger.Error(err)
			abort(c, http.StatusInternalServerError, err.Error())
		}

		return
	}

	models, total, err := s.service.ListProgramaProposta(programa.ID_PROGRAMA, pagination)
	switch err {
	case nil:
		c.Header("X-Total-Count", strconv.Itoa(total))
		c.JSON(http.StatusOK, models)
	default:
		s.logger.Error(err)
		abort(c, http.StatusInternalServerError, err.Error())
	}
}
