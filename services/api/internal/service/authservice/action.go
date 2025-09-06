package authservice

import (
	"context"

	"github.com/atareversei/quardian/services/api/internal/dto/authdto"
	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/authrepoparams"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
)

func (s *Service) GetActionId(ctx context.Context, req *authdto.GetActionIdRequest) (authdto.GetActionResponse, error) {
	const op = "authservice.GetActionId"

	action, err := s.repo.ReadActionIdByName(ctx, authrepoparams.ReadActionIdByNameInput{ActionName: req.ActionName})

	if err != nil {
		return authdto.GetActionResponse{}, richerror.
			New(op).
			WithError(err).
			WithKind(richerror.KindUnexpected)
	}

	return authdto.GetActionResponse{ActionId: action.ActionId}, nil
}
