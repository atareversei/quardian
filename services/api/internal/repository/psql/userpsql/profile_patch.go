package userpsql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/userrepoparams"
	"github.com/atareversei/quardian/services/api/pkg/errmsg"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
)

func (d *DB) EditProfilePartially(ctx context.Context, req userrepoparams.EditProfilePartiallyInput) (userrepoparams.EditProfilePartiallyOutput, error) {
	const op = "userpsql.EditProfilePartially"
	updates := map[string]interface{}{}

	if req.FirstName.HasValue() {
		updates["first_name"] = req.FirstName.Value
	}
	if req.LastName.HasValue() {
		updates["last_name"] = req.LastName.Value
	}
	if req.Username.HasValue() {
		updates["username"] = req.Username.Value
	}
	if req.Email.HasValue() {
		updates["email"] = req.Email.Value
	}
	if req.Mobile.HasValue() {
		updates["mobile"] = req.Mobile.Value
	}
	if req.BirthDate.HasValue() {
		updates["birth_date"] = req.BirthDate.Value
	}

	updates["status"] = "pending"

	query := "UPDATE users SET "
	params := []interface{}{}
	i := 1

	for k, v := range updates {
		query += fmt.Sprintf("%s = $%d,", k, i)
		params = append(params, v)
		i++
	}
	query = strings.TrimSuffix(query, ",")
	query += fmt.Sprintf(" WHERE user_id = $%d ", i)
	params = append(params, req.UserId)
	returning := `
	RETURNING
		user_id,
		employee_id,
		first_name,
		last_name,
		username,
		email,
		mobile,
		birth_date,
		status`
	query += returning

	row := d.db.Conn().QueryRowContext(ctx, query, params...)
	var output userrepoparams.EditProfilePartiallyOutput
	err := row.Scan(
		&output.UserId,
		&output.EmployeeId,
		&output.FirstName,
		&output.LastName,
		&output.Username,
		&output.Email,
		&output.Mobile,
		&output.BirthDate,
		&output.Status,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return output, richerror.New(op).WithKind(richerror.KindNotFound)
		}
		return output, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}

	return output, nil
}
