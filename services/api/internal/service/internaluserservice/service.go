package internaluserservice

import (
	"context"

	iurp "github.com/atareversei/quardian/services/api/internal/repository/repoparams/internaluserrepoparams"
)

type Repository interface {
	ListUsers(ctx context.Context, filters iurp.ListUsersInput) (iurp.ListUsersOutput, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}
