//	@title			Quardian API
//	@version		1.0
//	@description	Quardian services

//	@host		localhost:15340
//	@BasePath	/api/v1

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
package main

import (
	"flag"

	"github.com/atareversei/quardian/services/api/internal/delivery/httpserver/authhandler"
	"github.com/atareversei/quardian/services/api/internal/delivery/httpserver/internaluserhandler"
	"github.com/atareversei/quardian/services/api/internal/delivery/httpserver/userhandler"
	"github.com/atareversei/quardian/services/api/internal/repository/psql"
	"github.com/atareversei/quardian/services/api/internal/repository/psql/authpsql"
	"github.com/atareversei/quardian/services/api/internal/repository/psql/domains/internaluserpsql"
	"github.com/atareversei/quardian/services/api/internal/repository/psql/userpsql"
	"github.com/atareversei/quardian/services/api/internal/service/authservice"
	"github.com/atareversei/quardian/services/api/internal/service/internaluserservice"
	"github.com/atareversei/quardian/services/api/internal/service/userservice"
	"github.com/atareversei/quardian/services/api/internal/validator"
	"github.com/atareversei/quardian/services/api/internal/validator/authvalidator"
	"github.com/atareversei/quardian/services/api/internal/validator/internaluservalidator"
	"github.com/atareversei/quardian/services/api/internal/validator/uservalidator"
	"github.com/atareversei/quardian/services/api/pkg/authutil"
	"github.com/atareversei/quardian/services/api/pkg/project"

	"github.com/atareversei/quardian/services/api/internal/config"
	"github.com/atareversei/quardian/services/api/internal/delivery/httpserver"
	"github.com/atareversei/quardian/services/api/pkg/jwtutil"
	"github.com/atareversei/quardian/services/api/pkg/logger"
	"github.com/atareversei/quardian/services/api/pkg/translation"
)

func main() {
	env := flag.String("env", "dev", "specify env - `dev|test|prod`")
	flag.Parse()

	project.Init()
	cfg := config.Load(*env)
	logger.Init(cfg.Logger, *env)
	translation.Init(cfg.Language)
	jwtutil.Init(cfg.Auth.JWT)

	mainRepo := psql.New(cfg.Repository.Postgres)
	authRepo := authpsql.New(mainRepo)
	userRepo := userpsql.New(mainRepo)
	internalUserRepo := internaluserpsql.New(mainRepo)
	authutil.Init(authRepo)

	authSvc := authservice.New(authRepo)
	userSvc := userservice.New(userRepo)
	internalUserSvc := internaluserservice.New(internalUserRepo)

	mainValidator := validator.New()
	authValidator := authvalidator.New(mainValidator)
	userValidator := uservalidator.New(mainValidator)
	internalUserValidator := internaluservalidator.New(mainValidator)

	authHandler := authhandler.New(authValidator, authSvc)
	userHandler := userhandler.New(userValidator, userSvc)
	internalUserHandler := internaluserhandler.New(internalUserValidator, internalUserSvc)

	server := httpserver.New(httpserver.Args{
		AuthHandler:         *authHandler,
		UserHandler:         *userHandler,
		InternalUserHandler: *internalUserHandler,
		Config:              cfg,
	})

	server.Start()
}
