package echoutil

import (
	"net/http"

	"github.com/atareversei/quardian/services/api/pkg/contextutil"
	"github.com/atareversei/quardian/services/api/pkg/envelope"
	"github.com/atareversei/quardian/services/api/pkg/translation"
	"github.com/labstack/echo/v4"
)

func HandleEmptyPatchRequest(c echo.Context) error {
	ctx := c.Request().Context()
	lang := contextutil.GetLanguage(ctx)

	return c.JSON(http.StatusBadRequest, envelope.New(false).WithError(&envelope.ResponseError{
		Code:    envelope.ErrBadRequest,
		Message: translation.T(lang, "empty_patch_request"),
	}))
}
