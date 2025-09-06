package authdto

import "github.com/atareversei/quardian/services/api/internal/entity/authentity"

type GetResourceIdRequest struct {
	ResourceName authentity.ResourceName
}

type GetResourceResponse struct {
	ResourceId int
}
