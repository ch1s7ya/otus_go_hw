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

	for i, letter := range s {
		var err error

		count, err = strconv.Atoi(string(letter))
		// Current letter is letter
		if err != nil {
			if previousLetter == 0 {
				previousLetter = letter
				continue
			}

			if skip && letter != skipLetter {
				return "", ErrInvalidString
			}

			if skip && letter == skipLetter {
				previousLetter = skipLetter
				skip = false
				continue
			}

			if letter == skipLetter {
				skip = true
			}
			_, err := resultString.WriteRune(previousLetter)
			if err != nil {
				return "", ErrInvalidString
			}
			previousLetter = letter
			continue
		}

		// Current letter is number
		if i == 0 {
			return "", ErrInvalidString
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
