package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *server) RootHandler(c *gin.Context) {
	var routes []string
	for _, route := range s.router.Routes() {
		endpoint := fmt.Sprintf("%s: %s", route.Method, route.Path)
		routes = append(routes, endpoint)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "SICONV API Server",
		"routes":  routes,
	})
}

func (s *server) NotFoundHandler(c *gin.Context) {
	abort(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}
