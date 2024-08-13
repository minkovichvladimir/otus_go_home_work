package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(source string) (string, error) {
	if matched, err := regexp.MatchString(`^\d|([^\\]\d{2,})|\\[a-zA-Z]`, source); matched || err != nil {
		return "", ErrInvalidString
	}
	var b strings.Builder
	var prev string
	isEscapedBackslash := false
	for _, token := range source {
		if n, err := strconv.Atoi(string(token)); err == nil {
			if prev == `\` {
				if isEscapedBackslash {
					b.WriteString(strings.Repeat(prev, n))
					prev = ""
					isEscapedBackslash = false
				} else {
					prev = string(token)
				}
				continue
			}

			b.WriteString(strings.Repeat(prev, n))
			prev = ""
			continue
		}

		if prev == `\` {
			if isEscapedBackslash {
				b.WriteString(prev)
				isEscapedBackslash = false
			} else {
				isEscapedBackslash = true
			}
			prev = string(token)
			continue
		}

		if prev != "" {
			b.WriteString(prev)
		}
		prev = string(token)
	}
	b.WriteString(prev)
	return b.String(), nil
}
