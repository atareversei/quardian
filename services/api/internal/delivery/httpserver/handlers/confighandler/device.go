package confighandler

import (
	"github.com/atareversei/quardian/services/api/internal/dto/configdto"
	"github.com/atareversei/quardian/services/api/pkg/echoutil"
	"github.com/atareversei/quardian/services/api/pkg/envelope"
	"github.com/atareversei/quardian/services/api/pkg/queryparam"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h Handler) listDevices(c echo.Context) error {
	ctx := c.Request().Context()

	df := configdto.DeviceFilters{
		Page:    1,
		PerPage: 20,
	}
	if page, err := strconv.Atoi(c.QueryParam(queryparam.Page)); err == nil {
		df.Page = page
	}
	if perPage, err := strconv.Atoi(c.QueryParam(queryparam.PerPage)); err == nil {
		df.PerPage = perPage
	}
	req := configdto.ListDevicesRequest{DeviceFilters: df}

	validationErrors, err := h.Validator.ListDevices(ctx, req)
	if err != nil {
		return echoutil.HandleUnprocessableContent(c, validationErrors)
	}

	res, err := h.Service.ListDevices(ctx, req)
	if err != nil {
		return echoutil.HandleGenericError(c, err)
	}
	return c.JSON(http.StatusOK, envelope.New(true).WithData(res))
}
