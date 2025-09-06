package main

import (
	"flag"

	"github.com/atareversei/quardian/services/api/pkg/migrate"
)

func main() {
	// TODO: add limit flag
	up := flag.Bool("up", false, "If true migrate the database up, otherwise migrate it down")
	env := flag.String("env", "dev", "specify env - `dev|test|prod`")
	flag.Parse()

	migrate.Migrate(*up, 1, *env)
}
