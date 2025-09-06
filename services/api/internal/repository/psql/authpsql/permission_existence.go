package authpsql

import (
	"context"

	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/authrepoparams"
	"github.com/atareversei/quardian/services/api/pkg/errmsg"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
)

func (d *DB) IsUserIdPermittedOnResourceAndAction(ctx context.Context, input authrepoparams.IsUserIdPermittedOnResourceAndActionInput) (bool, error) {
	const op = "authpsql.IsUserIdPermittedOnResourceAndAction"

	const query = `
	SELECT EXISTS (
	  SELECT 1
	  FROM user_roles
		JOIN roles ON roles.role_id = user_roles.role_id
	  JOIN permissions ON user_roles.role_id = permissions.role_id
	  WHERE user_roles.user_id = $1
	    AND permissions.resource_id = $2
	    AND permissions.action_id = $3
	    AND roles.status = 'active'
	)`

	var exists bool
	err := d.db.Conn().QueryRowContext(ctx, query, input.UserId, input.ResourceId, input.ActionId).Scan(&exists)
	if err != nil {
		return false, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}
	return exists, nil
}
