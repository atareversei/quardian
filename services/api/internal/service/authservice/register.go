package authservice

import (
	"context"

	"github.com/atareversei/quardian/services/api/internal/dto/authdto"
	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/authrepoparams"
	"github.com/atareversei/quardian/services/api/pkg/contextutil"
	"github.com/atareversei/quardian/services/api/pkg/passwordhash"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
	"github.com/atareversei/quardian/services/api/pkg/translation"
)

func (s *Service) Register(ctx context.Context, req *authdto.RegisterRequest) (*authdto.RegisterResponse, error) {
	const op = "authservice.Register"
	lang := contextutil.GetLanguage(ctx)

	var (
		found       = false
		err   error = nil
	)

	found, err = s.repo.DoesUserNameExist(ctx, authrepoparams.DoesUserNmeExistInput{UserName: req.Username})

	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	if found {
		return nil, richerror.
			New(op).
			WithKind(richerror.KindInvalid).
			WithMessage(translation.T(lang, "username_already_exists"))
	}

	passwordHash, err := passwordhash.Hash(req.Password)

	if err != nil {
		return nil, richerror.
			New(op).
			WithError(err).
			WithKind(richerror.KindUnexpected)
	}

	_, err = s.repo.CreateUser(ctx, authrepoparams.CreateUserInput{UserName: req.Username, PasswordHash: passwordHash})

	if err != nil {
		return nil, richerror.
			New(op).
			WithKind(richerror.KindUnexpected)
	}

	return &authdto.RegisterResponse{}, nil
}
