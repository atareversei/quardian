package authrepoparams

import "github.com/atareversei/quardian/services/api/internal/entity/commonentity"

type ReadUserForLoginByUserNameInput struct {
	UserName string
}

type ReadUserForLoginByUserNameOutput struct {
	UserId       int
	UserName     string
	PasswordHash string
	Status       commonentity.Status
}
