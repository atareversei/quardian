package authservice

import (
	"context"

	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/authrepoparams"
)

type Repository interface {
	CreateUser(ctx context.Context, input authrepoparams.CreateUserInput) (authrepoparams.CreateUserOutput, error)
	DoesUserNameWithPasswordExist(ctx context.Context, input authrepoparams.DoesUserNameWithPasswordExistInput) (bool, error)
	DoesUserNameExist(ctx context.Context, input authrepoparams.DoesUserNmeExistInput) (bool, error)
	ReadUserForLoginByUserName(ctx context.Context, input authrepoparams.ReadUserForLoginByUserNameInput) (authrepoparams.ReadUserForLoginByUserNameOutput, bool, error)

	ReadResourceIdByName(ctx context.Context, input authrepoparams.ReadResourceIdByNameInput) (authrepoparams.ReadResourceByNameOutput, error)
	ReadActionIdByName(ctx context.Context, input authrepoparams.ReadActionIdByNameInput) (authrepoparams.ReadActionIdByNameOutput, error)
	IsUserIdPermittedOnResourceAndAction(ctx context.Context, input authrepoparams.IsUserIdPermittedOnResourceAndActionInput) (bool, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}
