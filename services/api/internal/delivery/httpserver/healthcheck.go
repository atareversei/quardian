package httpserver

import (
	"net/http"

	"github.com/atareversei/quardian/services/api/pkg/contextutil"
	"github.com/atareversei/quardian/services/api/pkg/translation"
	"github.com/labstack/echo/v4"
)

func (s *Server) healthCheck(c echo.Context) error {
	lang := contextutil.GetLanguage(c.Request().Context())

	return c.JSON(http.StatusOK, echo.Map{
		"message": translation.T(lang, "welcome"),
	})
}
