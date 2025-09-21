package confighandler

import (
	"github.com/atareversei/quardian/services/api/internal/service/configservice"
	"github.com/atareversei/quardian/services/api/internal/validator/configvalidator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service   configservice.Service
	Validator configvalidator.Validator
}

func New(service configservice.Service, validator configvalidator.Validator) Handler {
	return Handler{
		Service:   service,
		Validator: validator,
	}
}

func (h Handler) SetupRoutes(g *echo.Group) {
	configGroup := g.Group("/config")

	configGroup.GET("/devices", h.listDevices)
}
