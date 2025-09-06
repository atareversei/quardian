package middleware

import (
	"github.com/atareversei/quardian/services/api/pkg/logger"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// Logger returns a middleware that logs HTTP requests using zap logger
func Logger(env string) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:           true,
		LogStatus:        true,
		LogHost:          true,
		LogRemoteIP:      true,
		LogRequestID:     true,
		LogMethod:        true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogLatency:       true,
		LogError:         true,
		LogProtocol:      true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			// TODO: add config to check if logging is enabled
			errorMessage := ""
			operation := ""
			if v.Error != nil {
				re, ok := v.Error.(*richerror.RichError)
				if ok {
					errorMessage = re.GetMessage()
					operation = re.GetOperation()
				} else {
					errorMessage = v.Error.Error()
				}
			}

			mainFields := []zap.Field{
				zap.String("request_id", v.RequestID),
				zap.String("host", v.Host),
				zap.String("content_length", v.ContentLength),
				zap.String("protocol", v.Protocol),
				zap.String("method", v.Method),
				zap.Duration("latency", v.Latency),
				zap.String("error_message", errorMessage),
				zap.String("operation", operation),
				zap.String("remote_ip", v.RemoteIP),
				zap.Int64("response_size", v.ResponseSize),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
			}

			devFields := []zap.Field{
				zap.String("operation", operation),
				zap.String("uri", v.URI),
				zap.String("method", v.Method),
				zap.String("error_message", errorMessage),
			}

			var fields []zap.Field
			switch env {
			case "prod":
				fields = mainFields
			case "dev":
				fields = devFields
			}

			if v.Status >= 500 {
				logger.L().Named("http-server").Error("request", fields...)
			} else {
				logger.L().Named("http-server").Info("request", fields...)
			}

			return nil
		},
	})
}
