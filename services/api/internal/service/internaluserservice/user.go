package internaluserservice

import (
	"context"

	"github.com/atareversei/quardian/services/api/internal/dto/internaluserdto"
	"github.com/atareversei/quardian/services/api/internal/repository/repoparams/internaluserrepoparams"
	"github.com/atareversei/quardian/services/api/pkg/richerror"
)

func (s *Service) ListUsers(ctx context.Context, req *internaluserdto.ListUsersRequest) (*internaluserdto.ListUsersResponse, error) {
	const op = "internaluserservice.ListUsers"

	res, err := s.repo.ListUsers(ctx, internaluserrepoparams.ListUsersInput{Filters: req.UserFilters})
	if err != nil {
		return &internaluserdto.ListUsersResponse{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return &internaluserdto.ListUsersResponse{
		Meta: res.Meta,
		List: internaluserrepoparams.FromListUsersRepoParamToDTO(res.List),
	}, nil
}
