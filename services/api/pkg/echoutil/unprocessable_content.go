package echoutil

import (
	"net/http"

	"github.com/atareversei/quardian/services/api/internal/validator"
	"github.com/atareversei/quardian/services/api/pkg/contextutil"
	"github.com/atareversei/quardian/services/api/pkg/envelope"
	"github.com/atareversei/quardian/services/api/pkg/translation"
	"github.com/labstack/echo/v4"
)

func HandleUnprocessableContent(c echo.Context, validationErrors validator.ValidationErrors) error {
	ctx := c.Request().Context()
	lang := contextutil.GetLanguage(ctx)

	return c.JSON(http.StatusUnprocessableEntity, envelope.New(false).WithError(&envelope.ResponseError{
		Code:    envelope.ErrInvalidInput,
		Message: translation.T(lang, "invalid_input"),
		Fields:  validationErrors,
	}))
}
