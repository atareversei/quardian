package userpsql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/atareversei/quardian/services/api/internal/dto/userdto"
	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/userrepoparams"
	"github.com/atareversei/quardian/services/api/pkg/errmsg"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
)

func (d *DB) ReadProfileByUserId(ctx context.Context, input userrepoparams.ReadProfileByUserIdInput) (userrepoparams.ReadProfileByUserIdOutput, error) {
	const op = "userpsql.ReadProfileByUserId"

	row := d.db.Conn().QueryRowContext(ctx, `
	SELECT 
		user_id,
		employee_id,
		first_name,
		last_name,
		username,
		email,
		mobile,
		birth_date,
		status 
		FROM users
		WHERE user_id=$1`,
		input.UserId,
	)

	var output userrepoparams.ReadProfileByUserIdOutput
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
			return output, nil
		}
		return output, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}

	rows, err := d.db.Conn().QueryContext(ctx, `
	SELECT 
		permissions.role_id,
		permissions.resource_id,
		permissions.action_id
		FROM user_roles
		JOIN roles ON roles.role_id = user_roles.role_id
		JOIN permissions ON permissions.role_id = roles.role_id
		WHERE user_roles.user_id=$1 
		AND roles.status='active'`,
		input.UserId,
	)
	if err != nil {
		return userrepoparams.ReadProfileByUserIdOutput{}, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.QueryExecutionFailed).
			WithError(err)
	}

	defer rows.Close()

	grouped := make(map[int]map[int][]int)
	for rows.Next() {
		var roleId, resourceId, actionId int
		if err := rows.Scan(&roleId, &resourceId, &actionId); err != nil {
			return userrepoparams.ReadProfileByUserIdOutput{}, richerror.New(op).
				WithKind(richerror.KindUnexpected).
				WithMessage(errmsg.CantScanQueryResult).
				WithError(err)
		}

		if _, ok := grouped[roleId]; !ok {
			grouped[roleId] = make(map[int][]int)
		}
		if _, ok := grouped[roleId][resourceId]; !ok {
			grouped[roleId][resourceId] = make([]int, 1)
		}
		grouped[roleId][resourceId] = append(grouped[roleId][resourceId], actionId)
	}

	for roleID, permMap := range grouped {
		rawRole := userdto.RawRole{
			Id: roleID,
		}
		for permID, actionsSet := range permMap {
			rawPerm := userdto.RawPermission{
				Id: permID,
			}
			for action := range actionsSet {
				rawPerm.Actions = append(rawPerm.Actions, action)
			}
			rawRole.Permissions = append(rawRole.Permissions, rawPerm)
		}
		output.Roles = append(output.Roles, rawRole)
	}

	return output, nil
}
