package authservice

import (
	"context"

	"github.com/atareversei/quardian/services/api/internal/dto/authdto"
	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/authrepoparams"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
)

func (s *Service) IsUserIdPermittedOnResourceAndAction(ctx context.Context, req authdto.IsUserIdPermittedOnResourceAndActionRequest) (authdto.IsUserIdPermittedOnResourceAndActionResponse, error) {
	const op = "authservice.IsUserPermitted"
	isPermitted, err := s.repo.IsUserIdPermittedOnResourceAndAction(ctx, authrepoparams.IsUserIdPermittedOnResourceAndActionInput{UserId: req.UserId, ResourceId: req.ResourceId, ActionId: req.ActionId})
	if err != nil {
		return authdto.IsUserIdPermittedOnResourceAndActionResponse{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return authdto.IsUserIdPermittedOnResourceAndActionResponse{IsPermitted: isPermitted}, nil
}
