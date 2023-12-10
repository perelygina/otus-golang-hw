package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result strings.Builder
	var previous, current rune
	var isEscaped, backslashIsEscaped bool

	for _, current = range s {
		switch {
		case isEscaped:
			if previous == current {
				backslashIsEscaped = true
			}
			previous = current
			isEscaped = false
		case unicode.IsDigit(current):
			counter, _ := strconv.Atoi(string(current))
			if counter > 0 {
				result.WriteString(strings.Repeat(string(previous), counter))
			}
			if previous == 0 {
				return "", ErrInvalidString
			}
			previous = 0
		default:
			if previous != 0 {
				result.WriteRune(previous)
			}
			previous = current
		}

		if current == '\\' && !backslashIsEscaped {
			isEscaped = true
		} else if backslashIsEscaped {
			backslashIsEscaped = false
		}
	}

	if previous != 0 {
		result.WriteRune(previous)
	}

	return result.String(), nil
}
