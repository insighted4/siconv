package server

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
)

func (s *server) CreateEmpenhoDesembolsoHandler(c *gin.Context) {
	var empenhoDesembolso schema.EmpenhoDesembolso
	if err := c.BindJSON(&empenhoDesembolso); err != nil {
		s.logger.Error(err)
		abort(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := s.service.CreateEmpenhoDesembolso(&empenhoDesembolso)
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

	location := path.Join(Prefix, "empenho-desembolsos", id)
	c.Header("Location", location)
	c.Writer.WriteHeader(http.StatusCreated)
}
