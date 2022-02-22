package main

import (
	"math"
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

func Test_averageNumber(t *testing.T) {
	type testCase struct {
		sequence    string
		expected    float64
		expectedErr string
	}

	testCases := []testCase{
		{
			sequence:    "",
			expected:    0,
			expectedErr: ErrInvalidSequenceFormat.Error(),
		},
		{
			sequence:    "12-aa-",
			expected:    0,
			expectedErr: ErrInvalidSequenceFormat.Error(),
		},
		{
			sequence: "4-aa-8-bb",
			expected: 6,
		},
		{
			sequence: "5-aa-6-bb-10-cc",
			expected: 7,
		},
	}

	for _, tCase := range testCases {
		actual, err := averageNumber(tCase.sequence)
		if err != nil && err.Error() != tCase.expectedErr {
			t.Errorf("got err: %v, on sequence: %s", err, tCase.sequence)
		}

		if !equalsFloat(actual, tCase.expected) {
			t.Errorf("failed on sequence: %s, expected: %v, got: %v", tCase.sequence, tCase.expected, actual)
		}
	}
}

func equalsFloat(a, b float64) bool {
	return math.Abs(a-b) < 1e-6
}

func Test_wholeStory(t *testing.T) {
	type testCase struct {
		sequence    string
		expected    string
		expectedErr string
	}

	testCases := []testCase{
		{
			sequence:    "",
			expected:    "",
			expectedErr: ErrInvalidSequenceFormat.Error(),
		},
		{
			sequence:    "12-aa-",
			expected:    "",
			expectedErr: ErrInvalidSequenceFormat.Error(),
		},
		{
			sequence: "4-aa-8-bb",
			expected: "aa bb",
		},
		{
			sequence: "5-aa-6-bb-10-cc",
			expected: "aa bb cc",
		},
	}

	for _, tCase := range testCases {
		actual, err := wholeStory(tCase.sequence)
		if err != nil && err.Error() != tCase.expectedErr {
			t.Errorf("got err: %v, on sequence: %s", err, tCase.sequence)
		}

		if actual != tCase.expected {
			t.Errorf("failed on sequence: %s, expected: %v, got: %v", tCase.sequence, tCase.expected, actual)
		}
	}
}