package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const separator = "-"

var (
	validFmt                 = regexp.MustCompile(`(\d+-[[:ascii:]]+-)+$`)
	ErrInvalidSequenceFormat = errors.New("invalid sequence format")
)

// testValidity - validates input string.
// Expected a sequence of numbers followed by separator followed by text, eg: `23-ab-48-caba-56-haha`
// Time complexity: O(N)
// Estimated time: 20m
// Used time: 28m
func testValidity(s string) bool {
	if len(s) == 0 {
		return false
	}

	if s[len(s)-1:] == separator {
		return false
	}

	return validFmt.MatchString(s + separator)
}

// averageNumber - calculates average number from all the numbers
// Time complexity: O(N)
// Estimated time: 10m
// Used time: 17m
func averageNumber(s string) (float64, error) {
	if !testValidity(s) {
		return 0, ErrInvalidSequenceFormat
	}

	tokens := strings.Split(s, separator)

	sum := 0.0
	for i := 0; i < len(tokens); i += 2 {
		num, err := strconv.ParseUint(tokens[i], 10, 64)
		if err != nil {
			return 0, err
		}

		sum += float64(num)
	}

	n := float64(len(tokens) / 2)
	return sum / n, nil
}

// wholeStory - returns a text that is composed of all the text words separated by spaces
// Time complexity: O(N)
// Estimated time: 10m
// Used time: 9m
func wholeStory(s string) (string, error) {
	if !testValidity(s) {
		return "", ErrInvalidSequenceFormat
	}

	tokens := strings.Split(s, separator)
	result := make([]string, 0, len(tokens)/2)
	for i := 1; i < len(tokens); i += 2 {
		result = append(result, tokens[i])
	}

	return strings.Join(result, " "), nil
}
