package internaluserpsql

import (
	"context"
	"fmt"

	"github.com/atareversei/quardian/services/api/internal/dto"
	iurp "github.com/atareversei/quardian/services/api/internal/repository/repoparams/internaluserrepoparams"
	"github.com/atareversei/quardian/services/api/pkg/datetime"
	"github.com/atareversei/quardian/services/api/pkg/errmsg"
	"github.com/atareversei/quardian/services/api/pkg/repoutil"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
)

func (d *DB) ListUsers(ctx context.Context, input iurp.ListUsersInput) (iurp.ListUsersOutput, error) {
	const op = "internaluserpsql.ListUsers"
	const mainQuery = "SELECT user_id, employee_id, username, first_name, last_name, created_at, status FROM users WHERE 1=1 "
	const metaQuery = "SELECT COUNT(*) FROM users WHERE 1=1 "

	f := input.Filters
	ext := ""
	var args []interface{}
	argID := 1

	if f.Status != "" {
		ext += fmt.Sprintf(" AND status = $%d", argID)
		argID++
		args = append(args, f.Status)
	}

	var total int
	err := d.db.Conn().QueryRowContext(ctx, metaQuery+ext, args...).Scan(&total)
	if err != nil {
		return iurp.ListUsersOutput{}, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.QueryExecutionFailed).
			WithError(err)
	}

	ext += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argID, argID+1)
	argID += 2
	args = append(args, f.PerPage, (f.Page-1)*f.PerPage)

	rows, err := d.db.Conn().QueryContext(ctx, mainQuery+ext, args...)
	if err != nil {
		return iurp.ListUsersOutput{}, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.QueryExecutionFailed).
			WithError(err)
	}
	defer rows.Close()

	var users []iurp.ListUsersOutputItem
	for rows.Next() {
		var u iurp.ListUsersOutputItem
		if err := rows.Scan(
			&u.UserId,
			&u.EmployeeId,
			&u.UserName,
			&u.FirstName,
			&u.LastName,
			&u.CreatedAtRepo,
			&u.Status,
		); err != nil {
			return iurp.ListUsersOutput{}, richerror.New(op).
				WithKind(richerror.KindUnexpected).
				WithMessage(errmsg.CantScanQueryResult).
				WithError(err)
		}
		u.CreatedAt = datetime.ToStdDateTime(&u.CreatedAtRepo)
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return iurp.ListUsersOutput{}, fmt.Errorf("row iteration error: %w", err)
	}

	return iurp.ListUsersOutput{
		List: users,
		Meta: dto.ListMeta{
			CurrentPage:        input.Filters.Page,
			PerPage:            input.Filters.PerPage,
			LastPage:           repoutil.GetListLastPage(total, input.Filters.PerPage),
			TotalNumberOfItems: total,
		},
	}, nil
}
