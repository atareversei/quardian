package contextutil_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/atareversei/quardian/services/api/pkg/contextutil"
)

func TestGetUserId(t *testing.T) {
	type testCaseArg struct{ ctx context.Context }

	type testCaseExpected struct {
		userId int
		err    bool
	}

	type testCase struct {
		desc     string
		args     testCaseArg
		expected testCaseExpected
	}

	const userId = 1

	var testCases = []testCase{
		{
			desc:     "should return the user id if there is any available on the context",
			args:     testCaseArg{ctx: contextutil.WithUserId(context.Background(), strconv.Itoa(userId))},
			expected: testCaseExpected{userId: userId, err: true},
		},
		{
			desc:     "should return an error if there is no user id on the context",
			args:     testCaseArg{ctx: context.Background()},
			expected: testCaseExpected{userId: 0, err: false},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			userId, err := contextutil.GetUserID(tc.args.ctx)
			if tc.expected.err && err != nil {
				t.Errorf("Expected an error to be returned but instead function return nil")
			}
			if !tc.expected.err && tc.expected.userId != userId {
				t.Errorf("Expected: %d\nReceived: %d\n", tc.expected.userId, userId)
			}
		})
	}
}
