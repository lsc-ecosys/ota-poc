package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompareSemVers(t *testing.T) {
	var cases = []struct {
		desc     string
		input    [2]string
		expected int8
		err      error
	}{
		{
			desc:     "should compare two version strings correctly",
			input:    [2]string{"1.0.1", "1.0.0"},
			expected: int8(1),
			err:      nil,
		},
		{
			desc:     "should compare two version strings correctly",
			input:    [2]string{"1.0.0", "1.0.1"},
			expected: int8(-1),
			err:      nil,
		},
		{
			desc:     "should compare two version strings correctly",
			input:    [2]string{"1.0.1", "1.0.1"},
			expected: int8(2),
			err:      nil,
		},
	}

	for _, tc := range cases {
		output, err := CompareSemVers(tc.input[0], tc.input[1])
		require.Nil(t, err, fmt.Sprintf("got unexpected error: %s", err))
		if tc.err != nil {
			assert.Equal(t, tc.expected, output, tc.err)
		} else {
			assert.Equal(t, tc.expected, output)
		}
	}
}
