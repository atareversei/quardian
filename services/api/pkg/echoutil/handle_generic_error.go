package echoutil

import (
	"github.com/atareversei/quardian/services/api/pkg/envelope"
	"github.com/labstack/echo/v4"
)

func HandleGenericError(c echo.Context, err error) error {
	c.JSON(envelope.FromRichError(c, err))
	return err
}
