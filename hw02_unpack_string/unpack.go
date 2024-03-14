package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var resultString strings.Builder
	var previousLetter rune
	var count int
	skipLetter := '\\'
	var skip bool

	if len(s) == 0 {
		return "", nil
	}

	// Return err if last symbol equal `\`
	if s[len(s)-1] == byte(skipLetter) {
		return "", ErrInvalidString
	}

	for _, letter := range s {
		var err error

		count, err = strconv.Atoi(string(letter))
		// Current letter is letter
		if err != nil {
			switch {
			case previousLetter == 0:
				previousLetter = letter
				continue

			case skip && letter != skipLetter:
				return "", ErrInvalidString

			case skip && letter == skipLetter:
				previousLetter = skipLetter
				skip = false
				continue

			case letter == skipLetter:
				skip = true
			}

			_, err := resultString.WriteRune(previousLetter)
			if err != nil {
				return "", ErrInvalidString
			}

			previousLetter = letter
			continue
		}

		if previousLetter == 0 {
			return "", ErrInvalidString
		}

		if skip {
			previousLetter = letter
			skip = false
			continue
		}

		_, err = resultString.WriteString(strings.Repeat(string(previousLetter), count))
		if err != nil {
			return "", ErrInvalidString
		}

		previousLetter = 0
	}

	if previousLetter != 0 {
		_, err := resultString.WriteRune(previousLetter)
		if err != nil {
			return "", ErrInvalidString
		}
	}

	return resultString.String(), nil
}
