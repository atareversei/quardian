package userhandler

import (
	"net/http"

	"github.com/atareversei/quardian/services/api/internal/dto/userdto"
	"github.com/atareversei/quardian/services/api/pkg/echoutil"
	"github.com/atareversei/quardian/services/api/pkg/envelope"
	"github.com/atareversei/quardian/services/api/pkg/patch"
	"github.com/labstack/echo/v4"
)

func (h Handler) profile(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.service.Profile(ctx, &userdto.ProfileRequest{})
	if err != nil {
		return echoutil.HandleGenericError(c, err)
	}
	return c.JSON(http.StatusOK, envelope.New(true).WithData(res))
}

func (h Handler) editProfile(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(userdto.EditProfileRequest)
	if err := c.Bind(req); err != nil {
		return echoutil.HandleBadRequest(c)
	}
	isEmpty := patch.IsPatchStructEmpty(req)
	if isEmpty {
		return echoutil.HandleEmptyPatchRequest(c)
	}
	validationErrors, err := h.validator.EditProfile(ctx, req)
	if err != nil {
		return echoutil.HandleUnprocessableContent(c, validationErrors)
	}

	res, err := h.service.EditProfile(ctx, req)
	if err != nil {
		return echoutil.HandleGenericError(c, err)
	}
	return c.JSON(http.StatusOK, envelope.New(true).WithData(res))
}
