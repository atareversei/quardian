package authdto

import "github.com/atareversei/quardian/services/api/internal/entity/authentity"

type GetActionIdRequest struct {
	ActionName authentity.ActionName
}

type GetActionResponse struct {
	ActionId int
}
