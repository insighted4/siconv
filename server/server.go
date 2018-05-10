package server

import (
	"errors"

	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/insighted4/siconv/siconv"
	"github.com/insighted4/siconv/storage"
	"github.com/sirupsen/logrus"
)

const (
	Prefix = "/api/v1"
)

type Config struct {
	Token   string
	Storage storage.Service
	Logger  logrus.FieldLogger
}

type server struct {
	config  Config
	service siconv.Service
	logger  logrus.FieldLogger
	router  *gin.Engine
}

func (s *server) RunHTTPServer(addr string) error {
	s.logger.Infoln("Starting HTTP Server")
	return s.router.Run(addr)
}

func New(cfg Config) (*server, error) {
	if cfg.Storage == nil {
		return nil, errors.New("server: storage cannot be nil")
	}

	logger := cfg.Logger.WithField("component", "server")

	// Server
	srv := &server{
		config:  cfg,
		service: siconv.New(cfg.Storage, logger),
		router:  gin.New(),
		logger:  cfg.Logger.WithField("component", "server"),
	}

	r := srv.router

	if logger.Logger.Level >= logrus.InfoLevel {
		r.Use(gin.Logger())
	}
	r.Use(gin.Recovery())
	r.Use(LogMiddleware(cfg.Logger))
	r.Use(NewCORS())
	r.Use(RequestIDMiddleware())
	r.Use(VersionMiddleware())

	r.GET("/", srv.RootHandler)
	r.NoRoute(srv.NotFoundHandler)

	// API
	v1 := r.Group(Prefix)
	v1.GET("/consorcios", srv.ListConsorcioHandler)
	v1.GET("/consorcios/:id", srv.GetConsorcioHandler)

	v1.GET("/convenios", srv.ListConvenioHandler)
	v1.GET("/convenios/:id", srv.GetConvenioHandler)

	v1.GET("/desembolsos", srv.ListDesembolsoHandler)
	v1.GET("/desembolsos/:id", srv.GetDesembolsoHandler)

	v1.GET("/emendas", srv.ListEmendaHandler)
	v1.GET("/emendas/:id", srv.GetEmendaHandler)

	v1.GET("/empenhos", srv.ListEmpenhoHandler)
	v1.GET("/empenhos/:id", srv.GetEmpenhoHandler)

	v1.GET("/etapa-crono-fisicos", srv.ListEtapaCronoFisicoHandler)
	v1.GET("/etapa-crono-fisicos/:id", srv.GetEtapaCronoFisicoHandler)

	v1.GET("/historico-situacoes", srv.ListHistoricoSituacaoHandler)
	v1.GET("/historico-situacoes/:id", srv.GetHistoricoSituacaoHandler)

	v1.GET("/ingresso-contrapartidas", srv.ListIngressoContrapartidaHandler)
	v1.GET("/ingresso-contrapartidas/:id", srv.GetIngressoContrapartidaHandler)

	v1.GET("/meta-crono-fisicos", srv.ListMetaCronoFisicoHandler)
	v1.GET("/meta-crono-fisicos/:id", srv.GetMetaCronoFisicoHandler)

	v1.GET("/obtv-convenentes", srv.ListOBTVConvenenteHandler)
	v1.GET("/obtv-convenentes/:id", srv.GetOBTVConvenenteHandler)

	v1.GET("/pagamentos", srv.ListPagamentoHandler)
	v1.GET("/pagamentos/:id", srv.GetPagamentoHandler)

	v1.GET("/plano-aplicacao-detalhados", srv.ListPlanoAplicacaoDetalhadoHandler)
	v1.GET("/plano-aplicacao-detalhados/:id", srv.GetPlanoAplicacaoDetalhadoHandler)

	v1.GET("/programas", srv.ListProgramaHandler)
	v1.GET("/programas/:id", srv.GetProgramaHandler)

	v1.GET("/programa-propostas", srv.ListProgramaPropostaHandler)
	v1.GET("/programa-propostas/:id", srv.GetProgramaPropostaHandler)

	v1.GET("/proponentes", srv.ListProponenteHandler)
	v1.GET("/proponentes/:id", srv.GetProponenteHandler)

	v1.GET("/propostas", srv.ListPropostaHandler)
	v1.GET("/propostas/:id", srv.GetPropostaHandler)

	v1.GET("/prorroga-oficios", srv.ListProrrogaOficioHandler)
	v1.GET("/prorroga-oficios/:id", srv.GetProrrogaOficioHandler)

	v1.GET("/termo-aditivos", srv.ListTermoAditivoHandler)
	v1.GET("/termo-aditivos/:id", srv.GetTermoAditivoHandler)

	// Authenticated
	secured := r.Group(Prefix)
	secured.Use(AuthorizationMiddleware(cfg.Token))
	secured.POST("/consorcios", srv.CreateConsorcioHandler)
	secured.POST("/convenios", srv.CreateConvenioHandler)
	secured.POST("/desembolsos", srv.CreateDesembolsoHandler)
	secured.POST("/emendas", srv.CreateEmendaHandler)
	secured.POST("/empenhos", srv.CreateEmpenhoHandler)
	secured.POST("/empenho-desembolsos", srv.CreateEmpenhoDesembolsoHandler)
	secured.POST("/etapa-crono-fisicos", srv.CreateEtapaCronoFisicoHandler)
	secured.POST("/historico-situacoes", srv.CreateHistoricoSituacaoHandler)
	secured.POST("/ingresso-contrapartidas", srv.CreateIngressoContrapartidaHandler)
	secured.POST("/meta-crono-fisicos", srv.CreateMetaCronoFisicoHandler)
	secured.POST("/obtv-convenentes", srv.CreateOBTVConvenenteHandler)
	secured.POST("/pagamentos", srv.CreatePagamentoHandler)
	secured.POST("/plano-aplicacao-detalhados", srv.CreatePlanoAplicacaoDetalhadoHandler)
	secured.POST("/programas", srv.CreateProgramaHandler)
	secured.POST("/programa-propostas", srv.CreateProgramaPropostaHandler)
	secured.POST("/proponentes", srv.CreateProponenteHandler)
	secured.POST("/propostas", srv.CreatePropostaHandler)
	secured.POST("/prorroga-oficios", srv.CreateProrrogaOficioHandler)
	secured.POST("/termo-aditivos", srv.CreateTermoAditivoHandler)

	return srv, nil

}

func getPagination(c *gin.Context) *storage.Pagination {
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", strconv.Itoa(storage.Limit)))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	return storage.NewPagination(perPage, page)
}

func abort(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func NewCORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	allowHeaders := []string{
		"Accept",
		"Authorization",
		"Content-Type",
		"Keep-Alive",
		"Origin",
		"User-Agent",
		"X-Requested-With",
	}
	config.AllowHeaders = append(config.AllowHeaders, allowHeaders...)
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	return cors.New(config)
}
