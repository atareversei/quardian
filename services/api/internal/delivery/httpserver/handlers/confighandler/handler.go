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

func New(validator configvalidator.Validator, service configservice.Service) Handler {
	return Handler{
		Validator: validator,
		Service:   service,
	}
}

func (h Handler) SetupRoutes(g *echo.Group) {
	configGroup := g.Group("/config")

	configGroup.GET("/devices", h.listDevices)
}
