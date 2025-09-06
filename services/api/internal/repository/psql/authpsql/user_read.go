package authpsql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/authrepoparams"
	"github.com/atareversei/quardian/services/api/pkg/errmsg"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
)

func (d *DB) ReadUserForLoginByUserName(ctx context.Context, input authrepoparams.ReadUserForLoginByUserNameInput) (authrepoparams.ReadUserForLoginByUserNameOutput, bool, error) {
	const op = "authpsql.ReadUserForLoginByUserName"

	row := d.db.Conn().QueryRowContext(ctx, `SELECT user_id, username, password_hash, status FROM users WHERE username=$1`, input.UserName)

	var output authrepoparams.ReadUserForLoginByUserNameOutput
	err := row.Scan(&output.UserId, &output.UserName, &output.PasswordHash, &output.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return output, false, nil
		}
		return output, false, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}

	return output, true, nil
}
