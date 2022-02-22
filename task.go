package main

import (
	"regexp"
)

var validFmt = regexp.MustCompile(`(\d+-[[:ascii:]]+-)+$`)

// testValidity - validates input string.
// Expected a sequence of numbers followed by dash followed by text, eg: `23-ab-48-caba-56-haha`
// Time complexity: O(N)
// Estimated time: 20m
// Used time: 28m
func testValidity(s string) bool {
	if len(s) == 0 {
		return false
	}

	return validFmt.MatchString(s + "-")
}

// averageNumber - calculates average number from all the numbers
// Time complexity: O(N)
// Estimated time: 10m
// Used time: ?m
func averageNumber(s string) float64 {
	if len(s) == 0 {
		return 0
	}

	return 0
}
