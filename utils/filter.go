package utils

import "strings"

// lowerCase converts all strings in the input slice to lowercase.
func lowerCase(tokens []string) []string {

	r := make([]string, len(tokens))

	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}

	return r

}


func stopWordFilter(tokens []string) []string{





	
}
