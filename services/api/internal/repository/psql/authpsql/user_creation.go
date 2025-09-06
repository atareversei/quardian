package authpsql

import (
	"context"

	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/authrepoparams"
	"github.com/atareversei/quardian/services/api/pkg/errmsg"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
)

func (d *DB) CreateUser(ctx context.Context, input authrepoparams.CreateUserInput) (authrepoparams.CreateUserOutput, error) {
	const op = "authpsql.CreateUser"

	stmt, err := d.db.Conn().PrepareContext(ctx, `INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING user_id`)

	if err != nil {
		return authrepoparams.CreateUserOutput{}, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}

	row := stmt.QueryRowContext(ctx, input.UserName, input.PasswordHash)

	var output authrepoparams.CreateUserOutput
	err = row.Scan(&output.UserId)

	if err != nil {
		return authrepoparams.CreateUserOutput{}, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}
	return output, nil
}
