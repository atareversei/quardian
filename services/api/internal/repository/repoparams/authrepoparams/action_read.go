package authrepoparams

import (
	"github.com/atareversei/quardian/services/api/internal/entity/authentity"
	"github.com/atareversei/quardian/services/api/internal/entity/commonentity"
)

type ReadActionIdByNameInput struct {
	ActionName authentity.ActionName
}

type ReadActionIdByNameOutput struct {
	ActionId int
	Status   commonentity.Status
}
