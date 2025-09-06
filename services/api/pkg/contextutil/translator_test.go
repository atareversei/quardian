package contextutil_test

import (
	"context"
	"testing"

	"github.com/atareversei/quardian/services/api/pkg/contextutil"
	"github.com/atareversei/quardian/services/api/pkg/translation"
)

func TestWithLanguage(t *testing.T) {
	defaultLang := "fa"
	coreLang := "en"
	translation.Init(translation.Config{Default: defaultLang, Core: coreLang})

	type testCaseArg struct {
		lang string
	}

	type testCase struct {
		desc     string
		args     testCaseArg
		expected string
	}

	var testCases = []testCase{
		{
			desc:     "should return a context with the default language if provided with no language",
			args:     testCaseArg{lang: ""},
			expected: defaultLang,
		},
		{
			desc:     "should return a context with the language provided",
			args:     testCaseArg{lang: "sp"},
			expected: "sp",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := context.Background()
			actual := contextutil.WithLanguage(ctx, tc.args.lang)

			if actual.Value(contextutil.LanguageKey) != tc.expected {
				t.Errorf("Expected: %v\nReceived: %v", tc.expected, tc.args.lang)
			}
		})
	}
}

func TestGetLanguage(t *testing.T) {
	defaultLang := "fa"
	coreLang := "en"
	translation.Init(translation.Config{Default: defaultLang, Core: coreLang})

	type testCaseArg struct {
		ctx context.Context
	}

	type testCase struct {
		desc     string
		args     testCaseArg
		expected string
	}

	var testCases = []testCase{
		{
			desc:     "should return the language if the context has a language",
			args:     testCaseArg{ctx: contextutil.WithLanguage(context.Background(), "ar")},
			expected: "ar",
		},
		{
			desc:     "should return the default language if no language is available on the context",
			args:     testCaseArg{ctx: context.Background()},
			expected: defaultLang,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := contextutil.GetLanguage(tc.args.ctx)
			if actual != tc.expected {
				t.Errorf("Expected: %v\n Received: %v\n", tc.expected, actual)
			}
		})
	}
}
