package ustrings

import (
	"regexp"
	"strings"
	"sync"
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

var uppercaseAcronym = sync.Map{}

func ToCamel(s string, step string) string {
	ss := strings.Split(s, step)
	if len(ss) > 0 {
		for i, _s := range ss {
			ss[i] = toCamelInitCase(_s, false)
		}
	}

	return strings.Join(ss, ".")
}

func toCamelInitCase(s string, initCase bool) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	a, hasAcronym := uppercaseAcronym.Load(s)
	if hasAcronym {
		s = a.(string)
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := initCase
	prevIsCap := false
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		} else if prevIsCap && vIsCap && !hasAcronym {
			v += 'a'
			v -= 'A'
		}
		prevIsCap = vIsCap

		if vIsCap || vIsLow {
			n.WriteByte(v)
			capNext = false
		} else if vIsNum := v >= '0' && v <= '9'; vIsNum {
			n.WriteByte(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}
	return n.String()
}
