package authpsql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/authrepoparams"
	"github.com/atareversei/quardian/services/api/pkg/errmsg"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
)

func (d *DB) ReadResourceIdByName(ctx context.Context, input authrepoparams.ReadResourceIdByNameInput) (authrepoparams.ReadResourceByNameOutput, error) {
	const op = "authpsql.ReadResourceIdBasedOnName"

	row := d.db.Conn().QueryRowContext(ctx, `SELECT resource_id, status FROM resources WHERE name=$1`, input.ResourceName)

	var output authrepoparams.ReadResourceByNameOutput
	err := row.Scan(&output.ResourceId, &output.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return authrepoparams.ReadResourceByNameOutput{}, nil
		}
		return authrepoparams.ReadResourceByNameOutput{}, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}

	return output, nil
}
