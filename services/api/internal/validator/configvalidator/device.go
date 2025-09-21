package configvalidator

import (
	"context"
	"github.com/atareversei/quardian/services/api/internal/dto/configdto"
	"github.com/atareversei/quardian/services/api/internal/validator"
	"github.com/atareversei/quardian/services/api/pkg/contextutil"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const minPage = 1
const minPerPage = 5
const maxPerPage = 20

func (v *Validator) ListDevices(ctx context.Context, req configdto.ListDevicesRequest) (validator.ValidationErrors, error) {
	const op = "configvalidator.ListDevices"
	lang := contextutil.GetLanguage(ctx)

	if err := validation.ValidateStruct(req,
		validation.Field(&req.Page,
			validator.RequiredRule(lang, "fields.page"),
			validator.MinRule(lang, "fields.page", minPage),
		),
		validation.Field(&req.PerPage,
			validator.RequiredRule(lang, "fields.per_page"),
			validator.MinRule(lang, "fields.per_page", minPerPage),
			validator.MaxRule(lang, "fields.per_page", maxPerPage),
		),
	); err != nil {
		return v.util.Generate(validator.Args{
			Request:   req,
			Operation: op,
			Error:     err,
		})
	}
	return nil, nil
}
