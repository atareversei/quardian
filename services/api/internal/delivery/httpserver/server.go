package httpserver

import (
	"fmt"
	"path/filepath"

	"github.com/atareversei/quardian/services/api/internal/config"
	"github.com/atareversei/quardian/services/api/internal/delivery/httpserver/authhandler"
	"github.com/atareversei/quardian/services/api/internal/delivery/httpserver/internaluserhandler"
	"github.com/atareversei/quardian/services/api/internal/delivery/httpserver/middleware"
	"github.com/atareversei/quardian/services/api/internal/delivery/httpserver/userhandler"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

type Server struct {
	cfg                 config.Config
	Router              *echo.Echo
	authHandler         authhandler.Handler
	userHandler         userhandler.Handler
	internalUserHandler internaluserhandler.Handler
}

type Args struct {
	Config              config.Config
	AuthHandler         authhandler.Handler
	UserHandler         userhandler.Handler
	InternalUserHandler internaluserhandler.Handler
}

func New(args Args) *Server {
	return &Server{
		cfg:                 args.Config,
		Router:              echo.New(),
		authHandler:         args.AuthHandler,
		userHandler:         args.UserHandler,
		internalUserHandler: args.InternalUserHandler,
	}
}

func (s *Server) Start() {
	s.Router.Use(middleware.TranslatorMiddleware())
	s.Router.Use(middleware.Recovery())
	s.Router.Use(middleware.Logger(s.cfg.Env))
	// TODO: add CORS middleware to `middleware` package
	s.Router.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PATCH, echo.PUT, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	router := s.Router.Group("/api").Group("/v1")

	router.GET("/healthcheck", s.healthCheck)
	s.Router.Static("/docs", filepath.Join("docs", "api"))

	s.authHandler.SetRoutes(router)
	s.userHandler.SetRoutes(router)
	s.internalUserHandler.SetRoutes(router)

	addr := fmt.Sprintf(":%d", s.cfg.HttpServer.Port)

	if err := s.Router.Start(addr); err != nil {
		fmt.Println("router start error:", err)
	}
}
