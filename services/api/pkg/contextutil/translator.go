package contextutil

import (
	"context"

	"github.com/atareversei/quardian/services/api/pkg/translation"
)

func WithLanguage(ctx context.Context, lang string) context.Context {
	if lang == "" {
		lang = translation.GetDefaultLang()
	}
	return context.WithValue(ctx, LanguageKey, lang)
}

func GetLanguage(ctx context.Context) string {
	if lang, ok := ctx.Value(LanguageKey).(string); ok {
		return lang
	}

	return translation.GetDefaultLang()
}
