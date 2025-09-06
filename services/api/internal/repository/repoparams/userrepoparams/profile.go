package userrepoparams

import (
	"github.com/atareversei/quardian/services/api/internal/dto/userdto"
)

type ReadProfileByUserIdInput struct {
	UserId int
}

type ReadProfileByUserIdOutput struct {
	userdto.ProfileResponse
}

type EditProfilePartiallyInput struct {
	userdto.EditProfileRequest
	UserId int
}

type EditProfilePartiallyOutput struct {
	userdto.EditProfileResponse
}
