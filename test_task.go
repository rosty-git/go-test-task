package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// generate - generates random correct strings if the parameter is `true` and random incorrect strings otherwise
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

	return strings.Join(tokens, "-")
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
