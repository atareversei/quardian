package authhandler

import (
	"net/http"

	"github.com/atareversei/quardian/services/api/internal/dto/authdto"
	"github.com/atareversei/quardian/services/api/pkg/echoutil"
	"github.com/atareversei/quardian/services/api/pkg/envelope"
	"github.com/labstack/echo/v4"
)

// register let's a user sign up on the system with the minimum permissions
//
//	@Summary		User register
//	@Description	Sign up on system using a username and a password
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		authdto.RegisterRequest	true	"Register credentials"
//	@Success		200		{object}	envelope.OpenAPIResponseSuccess{data=authdto.RegisterResponse}
//	@Failure		400		{object}	envelope.OpenAPIResponseError
//	@Failure		422		{object}	envelope.OpenAPIResponseError
//	@Router			/users/register [post]
func (h Handler) register(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(authdto.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return echoutil.HandleBadRequest(c)
	}
	validationErrors, err := h.validator.Register(ctx, req)
	if err != nil {
		return echoutil.HandleUnprocessableContent(c, validationErrors)
	}
	res, err := h.service.Register(ctx, req)
	if err != nil {
		return echoutil.HandleGenericError(c, err)
	}
	return c.JSON(http.StatusOK, envelope.New(true).WithData(res))
}
