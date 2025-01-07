package tokenize

import (
	"searchSystem/internal/filter"
	"strings"
	"unicode"
)

func Tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func Analyze(text string) []string {
	tokens := Tokenize(text)
	tokens = filter.LowercaseFilter(tokens)
	tokens = filter.StemmerFilter(tokens)
	return tokens
}
