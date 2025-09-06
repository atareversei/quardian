package uservalidator

import (
	"context"
	"fmt"

	"github.com/atareversei/quardian/services/api/internal/dto/userdto"
	"github.com/atareversei/quardian/services/api/internal/validator"
	"github.com/atareversei/quardian/services/api/pkg/contextutil"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) EditProfile(ctx context.Context, req *userdto.EditProfileRequest) (validator.ValidationErrors, error) {
	const op = "uservalidator.EditProfile"
	lang := contextutil.GetLanguage(ctx)

	if err := validation.ValidateStruct(&req.Username,
		validation.Field(&req.Username.Value, validation.When(req.Username.IsSet, validator.UsernameRule(lang, "fields.username", true)...)),
	); err != nil {
		fmt.Println(err)
		return v.util.Generate(validator.Args{
			Request:   req,
			Operation: op,
			Error:     err,
		})
	}
	return nil, nil
}
