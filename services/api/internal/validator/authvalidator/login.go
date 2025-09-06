package authvalidator

import (
	"context"

	"github.com/atareversei/quardian/services/api/internal/dto/authdto"
	"github.com/atareversei/quardian/services/api/internal/validator"
	"github.com/atareversei/quardian/services/api/pkg/contextutil"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) Login(ctx context.Context, req *authdto.LoginRequest) (validator.ValidationErrors, error) {
	lang := contextutil.GetLanguage(ctx)
	const op = "uservalidator.Login"

	if err := validation.ValidateStruct(req,
		validation.Field(&req.Username,
			validator.RequiredRule(lang, "fields.username"),
			validator.LengthRule(lang, "fields.username", usernameMinLength, usernameMaxLength)),
		validation.Field(&req.Password, validator.PasswordRule(lang, "fields.password", true)...,
		)); err != nil {
		return v.util.Generate(validator.Args{
			Request:   req,
			Operation: op,
			Error:     err,
		})
	}
	return nil, nil
}
