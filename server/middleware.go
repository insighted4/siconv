package server

import (
	"fmt"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/insighted4/siconv/version"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func AuthorizationMiddleware(authorizationToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
		pieces := strings.Fields(c.GetHeader("Authorization"))
		if len(pieces) != 2 || authorizationToken != pieces[1] {
			abort(c, http.StatusUnauthorized, "invalid authorization token")
			return
		}

		c.Set("token", authorizationToken)
		c.Next()
	}
}

// LogMiddleware provide gin router handler.
func LogMiddleware(logger logrus.FieldLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := &logReq{
			URI:         c.Request.URL.Path,
			Method:      c.Request.Method,
			IP:          c.ClientIP(),
			ContentType: c.ContentType(),
			Agent:       c.Request.Header.Get("User-Agent"),
		}

		// format is string
		output := fmt.Sprintf("%s %s %s %s %s",
			log.Method,
			log.URI,
			log.IP,
			log.ContentType,
			log.Agent,
		)

		logger.Debug(output)
		c.Next()
	}
}

// RequestIDMiddleware injects a special header X-Request-Id to response headers
// that could be used to track incoming requests for monitoring/debugging
// purposes.
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
		c.Next()
	}
}

// VersionMiddleware : add version on header.
func VersionMiddleware() gin.HandlerFunc {
	// Set out header value for each response
	return func(c *gin.Context) {
		c.Header("X-Version", version.Version)
		c.Next()
	}
}
