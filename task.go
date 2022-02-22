package main

import (
	"errors"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const separator = "-"

var (
	validNumber              = regexp.MustCompile(`^\d+$`)
	validWord                = regexp.MustCompile(`^[[:ascii:]]+$`)
	ErrInvalidSequenceFormat = errors.New("invalid sequence format")
)

// testValidity - validates sequence string.
// Expected a sequence of numbers followed by separator followed by text, eg: `23-ab-48-caba-56-haha`
// Time complexity: O(N)
// Estimated time: 20m
// Used time: 28m
func testValidity(s string) bool {
	if len(s) == 0 {
		return false
	}

	tokens := strings.Split(s, separator)
	if len(tokens)%2 > 0 {
		return false
	}

	for i, token := range tokens {
		ok := false
		if i%2 == 0 {
			ok = validNumber.MatchString(token)
		} else {
			ok = validWord.MatchString(token)
		}
		if !ok {
			return false
		}
	}

	return true
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

// storyStats - returns four things:
//   * the shortest word
//   * the longest word
//   * the average word length
//   * the list (or empty list) of all words from the story that have the length the same as the average length rounded up and down.
// Time complexity: O(N)
// Estimated time: 20m
// Used time: 18m
func storyStats(s string) (shortestWord, longestWord string, avgWordLen float64, avgLenWords []string) {
	if !testValidity(s) {
		return
	}

	tokens := strings.Split(s, separator)

	sum := 0
	for i := 1; i < len(tokens); i += 2 {
		word := tokens[i]
		wordLen := len(tokens[i])
		sum += wordLen

		if shortestWord == "" || len(shortestWord) > wordLen {
			shortestWord = word
		}

		if longestWord == "" || len(longestWord) < wordLen {
			longestWord = word
		}
	}

	avgWordLen = float64(sum) / float64(len(tokens)/2)
	floor := math.Floor(avgWordLen)
	ceil := math.Ceil(avgWordLen)
	for i := 1; i < len(tokens); i += 2 {
		word := tokens[i]
		wordLen := float64(len(tokens[i]))
		if wordLen >= floor && wordLen <= ceil {
			avgLenWords = append(avgLenWords, word)
		}
	}

	return
}

// generate - generates random expected strings if the parameter is `true` and random incorrect strings otherwise
// Time complexity: O(N)
// Estimated time: 20m
// Used time: 30m
func generate(correct bool) string {
	rand.Seed(time.Now().UnixNano())

	if !correct && rand.Float32() < 0.1 {
		return ""
	}

	tokenPairs := rand.Intn(10) + 1
	tokens := make([]string, 0, tokenPairs*2)

	for i := 0; i < tokenPairs; i++ {
		if correct {
			tokens = append(tokens, genNumber())
		} else {
			tokens = append(tokens, genWord())
		}
		tokens = append(tokens, genWord())
	}

	return strings.Join(tokens, separator)
}

// genNumber - generates random number
func genNumber() string {
	return strconv.Itoa(rand.Intn(1000))
}

const asciiLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// genWord - generates random string
func genWord() string {
	wordLen := rand.Intn(15) + 1
	var word strings.Builder
	word.Grow(wordLen)

	for i := 0; i < wordLen; i++ {
		word.WriteByte(asciiLetters[rand.Intn(len(asciiLetters))])
	}

	return word.String()
}
