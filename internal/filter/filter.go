package filter

import (
	snowballeng "github.com/kljensen/snowball/english"
	"strings"
)

func LowercaseFilter(token []string) []string {
	r := make([]string, len(token))
	for i, token := range token {
		r[i] = strings.ToLower(token)
	}
	return r
}

func StemmerFilter(token []string) []string {
	r := make([]string, len(token))
	for i, token := range token {
		r[i] = snowballeng.Stem(token, false)
	}
	return r
}
