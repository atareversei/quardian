package contextutil

type contextKey string

const (
	TranslatorKey contextKey = "translator"
	LanguageKey   contextKey = "language"
	UserIdKey     contextKey = "user_id"
)
