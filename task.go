package main

import (
	"regexp"
)

// testValidity - validates input string.
// Expected a sequence of numbers followed by dash followed by text, eg: `23-ab-48-caba-56-haha`
// Time complexity: O(N)
// Estimated time: 20m
// Used time: 28m

var validFmt = regexp.MustCompile(`(\d+-[[:ascii:]]+-)+$`)

func testValidity(s string) bool {
	if len(s) == 0 {
		return false
	}

	return validFmt.MatchString(s + "-")
}
