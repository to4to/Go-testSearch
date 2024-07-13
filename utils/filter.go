package utils

import (
	"strings"

	snowballeng "github.com/kljensen/snowball/english"
)

// lowerCase converts all strings in the input slice to lowercase.
func lowerCase(tokens []string) []string {

	r := make([]string, len(tokens))

	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}

	return r

}

// stopWordFilter removes stop words from a given list of tokens.
// Stop words are common words that are often filtered out from text data.
// It takes a slice of strings representing tokens and returns a new slice with stop words removed.
func stopWordFilter(tokens []string) []string {

	stopWords := map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
		"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}

	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopWords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}

func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowballeng.Stem(token, false)
	}
	return r
}
