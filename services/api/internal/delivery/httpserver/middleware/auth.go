package middleware

import (
	"strconv"

	"github.com/atareversei/quardian/services/api/internal/dto/authdto"
	"github.com/atareversei/quardian/services/api/internal/entity/authentity"
	"github.com/atareversei/quardian/services/api/pkg/authutil"
	"github.com/atareversei/quardian/services/api/pkg/contextutil"
	"github.com/atareversei/quardian/services/api/pkg/envelope"
	"github.com/atareversei/quardian/services/api/pkg/jwtutil"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
	"github.com/labstack/echo/v4"
)

func IsAuthenticated() echo.MiddlewareFunc {
	const op = "middleware.isAuthenticated"

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()

			authorizationArray, ok := req.Header["Authorization"]
			if !ok {
				// TODO: add a message to unauthorized requests
				return c.JSON(envelope.FromRichError(c,
					richerror.
						New(op).
						WithKind(richerror.KindUnauthorized)))
			}

			authorization := authorizationArray[0][len("Bearer "):]
			userId, err := jwtutil.ValidateAndGetUserId(authorization)
			if err != nil {
				return c.JSON(envelope.FromRichError(c,
					richerror.
						New(op).
						WithKind(richerror.KindUnauthorized)))
			}
			ctx := contextutil.WithUserId(req.Context(), strconv.Itoa(userId))
			c.SetRequest(req.WithContext(ctx))

			return next(c)
		}
	}
}

func IsAuthorized(resource authentity.ResourceName, action authentity.ActionName) echo.MiddlewareFunc {
	const op = "middleware.IsAuthorized"
	authService := authutil.AuthService()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()

			r, err := authService.GetResourceId(req.Context(), &authdto.GetResourceIdRequest{ResourceName: resource})
			if err != nil {
				return c.JSON(envelope.FromRichError(c,
					richerror.
						New(op).
						WithKind(richerror.KindUnexpected).
						WithError(err)))
			}

			a, err := authService.GetActionId(req.Context(), &authdto.GetActionIdRequest{ActionName: action})
			if err != nil {
				return c.JSON(envelope.FromRichError(c,
					richerror.
						New(op).
						WithKind(richerror.KindUnexpected).
						WithError(err)))
			}

			userId, err := contextutil.GetUserID(req.Context())
			if err != nil {
				return c.JSON(envelope.FromRichError(c,
					richerror.
						New(op).
						WithKind(richerror.KindUnexpected).
						WithError(err)))
			}

			p, err := authService.IsUserIdPermittedOnResourceAndAction(
				req.Context(),
				authdto.IsUserIdPermittedOnResourceAndActionRequest{UserId: userId, ResourceId: r.ResourceId, ActionId: a.ActionId},
			)
			if err != nil {
				return c.JSON(envelope.FromRichError(c,
					richerror.
						New(op).
						WithKind(richerror.KindUnexpected).
						WithError(err)))
			}

			if !p.IsPermitted {
				return c.JSON(envelope.FromRichError(c,
					richerror.
						New(op).
						WithKind(richerror.KindForbidden)))
			}

			return next(c)
		}
	}
}
