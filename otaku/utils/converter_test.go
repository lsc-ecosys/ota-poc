package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStrToInt8(t *testing.T) {
	var cases = []struct {
		desc     string
		input    string
		expected int8
		err      error
	}{
		{
			desc:     "should convert string to int8 successfully",
			input:    "1",
			expected: int8(1),
			err:      nil,
		},
	}

	for _, tc := range cases {
		output, err := StrToInt8(tc.input)
		require.Nil(t, err, fmt.Sprintf("got unexpected error: %s", err))
		if tc.err != nil {
			assert.Equal(t, tc.expected, output, tc.err)
		} else {
			assert.Equal(t, tc.expected, output)
		}
	}
}
