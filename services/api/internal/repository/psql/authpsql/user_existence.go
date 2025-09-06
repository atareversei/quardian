package authpsql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/authrepoparams"
	"github.com/atareversei/quardian/services/api/pkg/errmsg"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
)

func (d *DB) DoesUserNameWithPasswordExist(ctx context.Context, input authrepoparams.DoesUserNameWithPasswordExistInput) (bool, error) {
	const op = "authpsql.DoesIdentifierWithPasswordExist"
	stmt, err := d.db.Conn().PrepareContext(ctx, "SELECT user_id FROM users WHERE username=$1 AND password_hash=$2")
	if err != nil {
		return false, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}

	row := stmt.QueryRowContext(ctx, input.UserName, input.PasswordHash)

	userId := new(int)
	err = row.Scan(userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}
	return true, nil
}

func (d *DB) DoesUserNameExist(ctx context.Context, input authrepoparams.DoesUserNmeExistInput) (bool, error) {
	const op = "authpsql.DoesUserNameExist"

	row := d.db.Conn().QueryRowContext(ctx, `SELECT user_id FROM users WHERE username=$1`, input.UserName)

	var userId = new(int)
	err := row.Scan(userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}

	return true, nil
}
