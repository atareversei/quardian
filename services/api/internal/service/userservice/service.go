package userservice

import (
	"context"

	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/userrepoparams"
)

type Repository interface {
	ReadProfileByUserId(ctx context.Context, input userrepoparams.ReadProfileByUserIdInput) (userrepoparams.ReadProfileByUserIdOutput, error)
	CanUserEditTheirProfile(ctx context.Context, userId int) bool
	EditProfilePartially(ctx context.Context, input userrepoparams.EditProfilePartiallyInput) (userrepoparams.EditProfilePartiallyOutput, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}
