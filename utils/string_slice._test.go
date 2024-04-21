package utils

import (
	"testing"
)

func TestStringOrSliceUnmarshalJSON(t *testing.T) {
	cases := []struct {
		input   string
		isError bool
	}{
		{
			input:   `"*"`,
			isError: false,
		},
		{
			input: `[
				"arn:aws:s3:::confidential-data",
				"arn:aws:s3:::confidential-data/*"
				]`,
			isError: false,
		},
		{
			input:   `{value: 10}`,
			isError: true,
		},
		{
			input: `[
				"arn:aws:s3:::confidential-data",
				[],
				"arn:aws:s3:::confidential-data/*"
				]`,
			isError: true,
		},
	}

	var stringOrSlice StringOrSlice
	for _, testCase := range cases {
		err := stringOrSlice.UnmarshalJSON([]byte(testCase.input))
		if (err == nil && testCase.isError) || (err != nil && !testCase.isError) {
			t.Fatalf("input: `%s`", testCase.input)
		}
	}
}
