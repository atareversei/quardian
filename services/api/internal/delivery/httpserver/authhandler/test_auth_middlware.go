package authhandler

import (
	"net/http"

	"github.com/atareversei/quardian/services/api/internal/dto/authdto"
	"github.com/atareversei/quardian/services/api/pkg/echoutil"
	"github.com/atareversei/quardian/services/api/pkg/envelope"
	"github.com/labstack/echo/v4"
)

func (h Handler) testAuthMiddleware(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(authdto.TestAuthMiddlewareRequest)
	if err := c.Bind(req); err != nil {
		return echoutil.HandleBadRequest(c)
	}
	res, err := h.service.TestAuthMiddleware(ctx, req)
	if err != nil {
		return echoutil.HandleGenericError(c, err)
	}
	return c.JSON(http.StatusOK, envelope.New(true).WithData(res))
}
