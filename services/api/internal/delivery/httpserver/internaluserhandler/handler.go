package internaluserhandler

import (
	"github.com/atareversei/quardian/services/api/internal/service/internaluserservice"
	"github.com/atareversei/quardian/services/api/internal/validator/internaluservalidator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	validator internaluservalidator.Validator
	service   internaluserservice.Service
}

func New(validator internaluservalidator.Validator, service internaluserservice.Service) *Handler {
	return &Handler{validator: validator, service: service}
}

func (h Handler) SetRoutes(g *echo.Group) {
	userGroup := g.Group("/internal/users")

	userGroup.GET("", h.listUsers)
}
