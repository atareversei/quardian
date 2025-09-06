package userservice

import (
	"context"

	"github.com/atareversei/quardian/services/api/internal/dto/userdto"
	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/userrepoparams"
	"github.com/atareversei/quardian/services/api/pkg/contextutil"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
	"github.com/atareversei/quardian/services/api/pkg/translation"
)

func (s *Service) Profile(ctx context.Context, req *userdto.ProfileRequest) (*userdto.ProfileResponse, error) {
	const op = "userservice.Profile"

	userId, err := contextutil.GetUserID(ctx)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindInvalid)
	}

	profile, err := s.repo.ReadProfileByUserId(ctx, userrepoparams.ReadProfileByUserIdInput{UserId: userId})
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return &profile.ProfileResponse, nil
}

func (s *Service) EditProfile(ctx context.Context, req *userdto.EditProfileRequest) (*userdto.EditProfileResponse, error) {
	const op = "userservice.EditProfile"
	lang := contextutil.GetLanguage(ctx)

	userId, err := contextutil.GetUserID(ctx)
	if err != nil {
		return &userdto.EditProfileResponse{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	canEdit := s.repo.CanUserEditTheirProfile(ctx, userId)
	if !canEdit {
		return &userdto.EditProfileResponse{}, richerror.New(op).WithKind(richerror.KindForbidden).WithMessage(translation.T(lang, "user_not_active"))
	}

	newProfile, err := s.repo.EditProfilePartially(ctx, userrepoparams.EditProfilePartiallyInput{EditProfileRequest: *req, UserId: userId})
	if err != nil {
		return &userdto.EditProfileResponse{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return &userdto.EditProfileResponse{ProfileResponse: newProfile.ProfileResponse}, nil
}
