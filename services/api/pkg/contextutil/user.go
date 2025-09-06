package contextutil

import (
	"context"
	"strconv"

	"github.com/atareversei/quardian/services/api/pkg/richerror"
)

func WithUserId(ctx context.Context, userId string) context.Context {
	return context.WithValue(ctx, UserIdKey, userId)
}

func GetUserID(ctx context.Context) (int, error) {
	const op = "contextutil.GetUserId"

	userIdStr, ok := ctx.Value(UserIdKey).(string)

	if !ok {
		return 0, richerror.New(op).WithMessage("`user_id` field is not present on context")
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return 0, richerror.New(op).WithError(err)
	}

	return userId, nil
}
