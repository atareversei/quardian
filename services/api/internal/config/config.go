package config

import (
	"github.com/atareversei/quardian/services/api/internal/repository/psql"
	"github.com/atareversei/quardian/services/api/pkg/jwtutil"
	"github.com/atareversei/quardian/services/api/pkg/logger"
	"github.com/atareversei/quardian/services/api/pkg/translation"
)

type HTTPServerConfig struct {
	Port uint `koanf:"port"`
}

type RepositoryConfig struct {
	Postgres psql.Config `koanf:"postgres"`
}

type AuthType string

const (
	AuthTypeRefreshAndAccess AuthType = "refresh_and_access"
	AuthTypeRegularToken     AuthType = "regular_token"
)

type AuthConfig struct {
	AuthType AuthType          `koanf:"auth_type"`
	JWT      jwtutil.JWTConfig `koanf:"jwt"`
}

type Config struct {
	Env        string             `koanf:"env"`
	HttpServer HTTPServerConfig   `koanf:"http_server"`
	Repository RepositoryConfig   `koanf:"repository"`
	Language   translation.Config `koanf:"language"`
	Logger     logger.Config      `koanf:"logger"`
	Auth       AuthConfig         `koanf:"auth"`
}
