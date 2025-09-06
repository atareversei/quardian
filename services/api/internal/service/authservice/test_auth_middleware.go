package authservice

import (
	"context"

	"github.com/atareversei/quardian/services/api/internal/dto/authdto"
)

func (s *Service) TestAuthMiddleware(ctx context.Context, req *authdto.TestAuthMiddlewareRequest) (*authdto.TestAuthMiddlewareResponse, error) {
	return &authdto.TestAuthMiddlewareResponse{}, nil
}
