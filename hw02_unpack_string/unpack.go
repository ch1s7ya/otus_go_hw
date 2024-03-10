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
	var skipLetter = '\\'
	var skip bool

	if len(s) == 0 {
		return "", nil
	}

	if _, err := strconv.Atoi(string(s[0])); err == nil {
		return "", ErrInvalidString
	}

	for _, letter := range s {
		var err error

		if previousLetter == 0 && count != 0 {
			_, err = strconv.Atoi(string(letter))
			if err == nil {
				return "", ErrInvalidString
			}
		}

		if previousLetter == 0 {
			_, err = strconv.Atoi(string(letter))
			if err != nil {
				previousLetter = letter
				continue
			}
			return "", ErrInvalidString
		}

		count, err = strconv.Atoi(string(letter))
		// Current letter is letter
		if err != nil {
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
			count = 0
			continue
		}

		// Current letter is number
		if skip {
			previousLetter = letter
			skip = false
			count = 0
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
