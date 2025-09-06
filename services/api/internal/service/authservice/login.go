package authservice

import (
	"context"

	"github.com/atareversei/quardian/services/api/internal/dto/authdto"
	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/authrepoparams"
	"github.com/atareversei/quardian/services/api/pkg/contextutil"
	"github.com/atareversei/quardian/services/api/pkg/jwtutil"
	"github.com/atareversei/quardian/services/api/pkg/passwordhash"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
	"github.com/atareversei/quardian/services/api/pkg/translation"
)

func (s *Service) Login(ctx context.Context, req *authdto.LoginRequest) (*authdto.LoginResponse, error) {
	const op = "authservice.Login"
	lang := contextutil.GetLanguage(ctx)

	var (
		user  authrepoparams.ReadUserForLoginByUserNameOutput
		found       = false
		err   error = nil
	)

	user, found, err = s.repo.ReadUserForLoginByUserName(ctx, authrepoparams.ReadUserForLoginByUserNameInput{UserName: req.Username})

	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	if !found {
		return nil, richerror.
			New(op).
			WithKind(richerror.KindNotFound).
			WithMessage(translation.T(lang, "user_not_found"))
	}

	areIdentical := passwordhash.Compare(user.PasswordHash, req.Password)

	if !areIdentical {
		return nil, richerror.
			New(op).
			WithKind(richerror.KindNotFound).
			WithMessage(translation.T(lang, "user_not_found"))
	}

	token, err := jwtutil.Create(user.UserId)

	if err != nil {
		return nil, richerror.
			New(op).
			WithKind(richerror.KindUnexpected).
			WithMeta(map[string]any{
				"username": user.UserName,
			})
	}

	return &authdto.LoginResponse{Token: token}, nil
}
