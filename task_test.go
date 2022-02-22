package main

import (
	"testing"
)

func Test_testValidity(t *testing.T) {
	type testCase struct {
		sequence string
		expected bool
	}

	testCases := []testCase{
		{
			sequence: "",
			expected: false,
		},
		{
			sequence: "1",
			expected: false,
		},
		{
			sequence: "1-",
			expected: false,
		},
		{
			sequence: "1-a-",
			expected: false,
		},
		{
			sequence: "1-a-a",
			expected: false,
		},
		{
			sequence: "1-a-1",
			expected: false,
		},
		{
			sequence: "1-a-1-",
			expected: false,
		},
		{
			sequence: "1-a-1--",
			expected: false,
		},
		{
			sequence: "1-a-1--a",
			expected: false,
		},
		{
			sequence: generate(true),
			expected: true,
		},
		{
			sequence: generate(true),
			expected: true,
		},
		{
			sequence: generate(false),
			expected: false,
		},
		{
			sequence: generate(false),
			expected: false,
		},
	}

	for _, tCase := range testCases {
		actual := testValidity(tCase.sequence)
		if actual != tCase.expected {
			t.Errorf("failed on sequence: %s, expected: %v", tCase.sequence, tCase.expected)
		}
	}
}
