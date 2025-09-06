package authpsql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/authrepoparams"
	"github.com/atareversei/quardian/services/api/pkg/errmsg"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
)

func (d *DB) ReadActionIdByName(ctx context.Context, input authrepoparams.ReadActionIdByNameInput) (authrepoparams.ReadActionIdByNameOutput, error) {
	const op = "authpsql.ReadActionIdByName"

	row := d.db.Conn().QueryRowContext(ctx, `SELECT action_id, status FROM actions WHERE name=$1`, input.ActionName)

	var output authrepoparams.ReadActionIdByNameOutput
	err := row.Scan(&output.ActionId, &output.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return authrepoparams.ReadActionIdByNameOutput{}, nil
		}
		return authrepoparams.ReadActionIdByNameOutput{}, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}

	return output, nil
}
