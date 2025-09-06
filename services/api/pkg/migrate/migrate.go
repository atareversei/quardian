package migrate

import (
	"github.com/atareversei/quardian/services/api/internal/config"
	"github.com/atareversei/quardian/services/api/internal/repository/psql/pgmigrator"
)

// TODO: add limit
func Migrate(up bool, limit int, env string) {
	cfg := config.Load(env)
	migrator := pgmigrator.NewMigrator(cfg.Repository.Postgres)

	if up {
		migrator.Up()
	} else {
		migrator.Down()
	}
}
