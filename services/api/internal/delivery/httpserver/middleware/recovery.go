package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/atareversei/quardian/services/api/pkg/envelope"
	"github.com/atareversei/quardian/services/api/pkg/logger"
	"github.com/atareversei/quardian/services/api/pkg/translation"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Recovery() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			lang := c.Request().Header.Get("Accept-Language")

			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}

					stack := debug.Stack()

					fields := []zap.Field{
						zap.String("panic", err.Error()),
						zap.String("stack", string(stack)),
						zap.String("method", c.Request().Method),
						zap.String("path", c.Request().URL.Path),
						zap.String("remote_ip", c.RealIP()),
					}

					logger.L().Named("panic").Error("recovered from panic", fields...)

					c.JSON(http.StatusInternalServerError, envelope.New(false).WithError(&envelope.ResponseError{
						Code:    envelope.ErrInternal,
						Message: translation.T(lang, "internal_server"),
					}))
				}
			}()

			return next(c)
		}
	}
}
