package ustrings

import (
	"regexp"
	"strings"
	"unicode"
)

func ToPascalCase(input string) string {
	re := regexp.MustCompile(`[\W_]+`)
	words := re.Split(input, -1)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.Title(word)
		}
	}
	return strings.Join(words, "")
}

func PascalToCamel(s string) string {
	if s == "" {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])

	return string(r)
}

func ToSnake(s string) string {
	var result []rune
	if strings.Contains(s, " ") {
		s = strings.ReplaceAll(s, " ", "")
	}

	for i, r := range s {
		if unicode.IsUpper(r) {
			if i != 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func PascalToSnake(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i != 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
