package config

import (
	"path/filepath"
	"time"

	"github.com/atareversei/quardian/services/api/pkg/project"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// TODO: turn environment into a type
func Load(environment string) Config {
	yamlCfgPath := filepath.Join(project.GetProjectRoot(), "config.yml")

	cfg := Config{
		Env: environment,
	}
	k := koanf.New(".")
	err := k.Load(confmap.Provider(defaultConfig, "."), nil)
	if err != nil {
		panic("failed to initialize the default config")
	}

	err = k.Load(file.Provider(yamlCfgPath), yaml.Parser())
	if err != nil {
		panic("failed to read configurations from config.yml")
	}

	var dotenvFileName = ""
	switch environment {
	case "test":
		dotenvFileName = ".env.test"
	default:
		dotenvFileName = ".env.dev"
	}

	dotenv := NewEnv(EnvPrefix, dotenvFileName)
	dotenv.Load()
	err = k.Load(confmap.Provider(map[string]any{
		"repository.postgres.username":    dotenv.Get("POSTGRES_USER"),
		"repository.postgres.password":    dotenv.Get("POSTGRES_PASSWORD"),
		"repository.postgres.host":        dotenv.Get("POSTGRES_HOST"),
		"repository.postgres.port":        dotenv.Get("POSTGRES_PORT"),
		"repository.postgres.dbname":      dotenv.Get("POSTGRES_DB"),
		"auth.jwt.basic_token_secret_key": dotenv.Get("JWT_BASIC_TOKEN_SECRET_KEY"),
		// TODO: investigate basic_token_expiration_time_ns -- why ns
		"auth.jwt.basic_token_expiration_time_days": time.Duration(dotenv.GetNumber("JWT_BASIC_TOKEN_EXPIRATION_TIME_DAYS")),
	}, "."), nil)

	if err != nil {
		panic("failed to load the env config")
	}

	err = k.Unmarshal("", &cfg)

	if err != nil {
		panic("failed to unmarshal the config")
	}

	return cfg
}
