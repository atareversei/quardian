package authservice

import (
	"context"

	"github.com/atareversei/quardian/services/api/internal/dto/authdto"
	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/authrepoparams"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
)

func (s *Service) GetResourceId(ctx context.Context, req *authdto.GetResourceIdRequest) (authdto.GetResourceResponse, error) {
	const op = "authservice.GetResourceId"

	resource, err := s.repo.ReadResourceIdByName(ctx, authrepoparams.ReadResourceIdByNameInput{ResourceName: req.ResourceName})

	if err != nil {
		return authdto.GetResourceResponse{}, richerror.
			New(op).
			WithError(err).
			WithKind(richerror.KindUnexpected)
	}

	return authdto.GetResourceResponse{ResourceId: resource.ResourceId}, nil
}
