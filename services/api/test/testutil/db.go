package testutil

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/atareversei/quardian/services/api/internal/config"
	"github.com/atareversei/quardian/services/api/internal/repository/psql"
	"github.com/atareversei/quardian/services/api/pkg/migrate"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

type TestDB struct {
	DB       *sql.DB
	Resource *dockertest.Resource
	Pool     *dockertest.Pool
}

func SetupTestDB() *TestDB {
	cfg := config.Load("test")

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "16.4", // TODO: add to config
		Env: []string{
			fmt.Sprintf("POSTGRES_USER=%s", cfg.Repository.Postgres.Username),
			fmt.Sprintf("POSTGRES_PASSWORD=%s", cfg.Repository.Postgres.Password),
			fmt.Sprintf("POSTGRES_DB=%s", cfg.Repository.Postgres.DBName),
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
	})
	if err != nil {
		log.Fatalf("Could not start dockertest resource: %s", err)
	}

	dsn := psql.GetDSN(cfg.Repository.Postgres)
	var db *sql.DB
	pool.Retry(func() error {
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			return err
		}
		return db.Ping()
	})

	migrate.Migrate(true, 1, "test")

	// seed data

	return &TestDB{
		DB:       db,
		Resource: resource,
		Pool:     pool,
	}
}
