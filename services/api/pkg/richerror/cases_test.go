package richerror

type testCase struct {
	description string
	args        RichError
	expected    any
}

var getOperationTestCases = []testCase{
	{
		description: "should return the correct operation value",
		args: RichError{
			operation: "Operation1",
		},
		expected: "Operation1",
	},
	{
		description: "should return the correct operation value",
		args: RichError{
			operation: "Operation1",
			wrappedError: &RichError{
				operation: "Operation2",
			},
		},
		expected: "Operation2\n\tOperation1",
	},
	{
		description: "should return the correct operation value",
		args: RichError{
			operation: "Operation1",
			wrappedError: &RichError{
				operation: "Operation2",
				wrappedError: &RichError{
					operation: "Operation3",
				},
			},
		},
		expected: "Operation3\n\tOperation2\n\tOperation1",
	},
}

var getKindTestCases = []testCase{
	{
		description: "should return the correct kind value",
		args: RichError{
			operation: "Operation2",
			kind:      KindForbidden,
		},
		expected: KindForbidden,
	},
	{
		description: "should return the correct kind value from the nested error",
		args: RichError{
			operation: "Operation2",
			wrappedError: &RichError{
				operation: "Operation3",
				kind:      KindInvalid,
			},
		},
		expected: KindInvalid,
	},
}

var getMessageTestCases = []testCase{
	{
		description: "should return the correct error message",
		args: RichError{
			operation: "Operation3",
			message:   "some error message",
		},
		expected: "some error message",
	},
	{
		description: "should return the correct error message from the nested error",
		args: RichError{
			operation: "Operation3",
			wrappedError: &RichError{
				operation: "Operation4",
				message:   "nested error message",
			},
		},
		expected: "nested error message",
	},
	{
		description: "should return the correct error message with details",
		args: RichError{
			operation: "Operation3",
			message:   "parent error message",
			wrappedError: &RichError{
				operation: "Operation4",
				message:   "nested error message",
			},
		},
		expected: "parent error message More Information: nested error message",
	},
}

var getErrorTestCases = append(getMessageTestCases, []testCase{
	{
		description: "should return the 'unspecified error'",
		args: RichError{
			operation: "Operation6",
		},
		expected: "unspecified error",
	},
	{
		description: "should return the 'unspecified error'",
		args: RichError{
			operation: "Operation6",
			wrappedError: &RichError{
				operation: "Operation7",
			},
		},
		expected: "unspecified error",
	}, {
		description: "should return the 'unspecified error'",
		args: RichError{
			operation: "Operation6",
			wrappedError: &RichError{
				operation: "Operation7",
				message:   "nested error message",
			},
		},
		expected: "nested error message",
	},
}...)

var getMetaTestCases = []testCase{
	{
		description: "should return the correct meta",
		args: RichError{
			operation: "Operation5",
			message:   "Some message",
			meta: map[string]any{
				"name": "NameForMeta",
			},
		},
		expected: map[string]any{
			"name": "NameForMeta",
		},
	},
}
