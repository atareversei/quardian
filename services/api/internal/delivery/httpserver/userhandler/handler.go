package userhandler

import (
	"github.com/atareversei/quardian/services/api/internal/delivery/httpserver/middleware"
	"github.com/atareversei/quardian/services/api/internal/entity/authentity"
	"github.com/atareversei/quardian/services/api/internal/service/userservice"
	"github.com/atareversei/quardian/services/api/internal/validator/uservalidator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	validator uservalidator.Validator
	service   userservice.Service
}

func New(validator uservalidator.Validator, service userservice.Service) *Handler {
	return &Handler{validator: validator, service: service}
}

func (h Handler) SetRoutes(g *echo.Group) {
	authGroup := g.Group("/users", middleware.IsAuthenticated())

	authGroup.GET("/profile", h.profile, middleware.IsAuthorized(authentity.ResourceUserProfile, authentity.ActionRead))
	authGroup.PATCH("/profile", h.editProfile)
	// TODO:  middleware.IsAuthorized(authentity.ResourceUserProfile, authentity.ActionPatch)
}
