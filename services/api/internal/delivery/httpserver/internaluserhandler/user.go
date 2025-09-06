package internaluserhandler

import (
	"net/http"
	"strconv"

	"github.com/atareversei/quardian/services/api/internal/dto/internaluserdto"
	"github.com/atareversei/quardian/services/api/pkg/echoutil"
	"github.com/atareversei/quardian/services/api/pkg/envelope"
	"github.com/atareversei/quardian/services/api/pkg/queryparam"
	"github.com/labstack/echo/v4"
)

func (h Handler) listUsers(c echo.Context) error {
	ctx := c.Request().Context()

	uf := internaluserdto.UserFilters{
		Page:    1,
		PerPage: 50,
	}

	// TODO: validation
	if page, err := strconv.Atoi(c.QueryParam(queryparam.Page)); err == nil {
		uf.Page = page
	}
	if perPage, err := strconv.Atoi(c.QueryParam(queryparam.PerPage)); err == nil {
		uf.PerPage = perPage
	}

	res, err := h.service.ListUsers(ctx, &internaluserdto.ListUsersRequest{UserFilters: uf})
	if err != nil {
		return echoutil.HandleGenericError(c, err)
	}
	return c.JSON(http.StatusOK, envelope.New(true).WithData(res))
}
