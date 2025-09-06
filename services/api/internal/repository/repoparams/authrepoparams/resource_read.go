package authrepoparams

import (
	"github.com/atareversei/quardian/services/api/internal/entity/authentity"
	"github.com/atareversei/quardian/services/api/internal/entity/commonentity"
)

type ReadResourceIdByNameInput struct {
	ResourceName authentity.ResourceName
}

type ReadResourceByNameOutput struct {
	ResourceId int
	Status     commonentity.Status
}
