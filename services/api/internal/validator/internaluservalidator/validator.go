package internaluservalidator

import "github.com/atareversei/quardian/services/api/internal/validator"

type Validator struct {
	util *validator.Validator
}

func New(util *validator.Validator) Validator {
	return Validator{
		util: util,
	}
}
