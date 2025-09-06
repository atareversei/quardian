package authhandler

import (
	"github.com/atareversei/quardian/services/api/internal/delivery/httpserver/middleware"
	"github.com/atareversei/quardian/services/api/internal/entity/authentity"
	"github.com/atareversei/quardian/services/api/internal/service/authservice"
	"github.com/atareversei/quardian/services/api/internal/validator/authvalidator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	validator authvalidator.Validator
	service   authservice.Service
}

func New(validator authvalidator.Validator, service authservice.Service) *Handler {
	return &Handler{validator: validator, service: service}
}

func (h Handler) SetRoutes(g *echo.Group) {
	authGroup := g.Group("/auth")

	authGroup.POST("/register", h.register)
	authGroup.POST("/login", h.login)
	authGroup.GET("/test-auth-middleware", h.testAuthMiddleware, middleware.IsAuthorized(authentity.ResourceAuth, authentity.ActionRead))
}
