package userpsql

import (
	"context"
)

func (d *DB) CanUserEditTheirProfile(ctx context.Context, userId int) bool {

	var status string
	row := d.db.Conn().QueryRowContext(ctx, `
	SELECT 
		status
		FROM users
		WHERE user_id=$1`,
		userId,
	)
	row.Scan(&status)
	return status == "active"
}
