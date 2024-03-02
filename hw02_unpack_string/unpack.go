package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var resultString strings.Builder
	var currentLetter rune
	var count int

	if len(s) == 0 {
		return "", nil
	}

	if _, err := strconv.Atoi(string(s[0])); err == nil {
		return "", ErrInvalidString
	}

	for _, letter := range s {
		var err error
		if currentLetter == 0 && count != 0 {
			_, err = strconv.Atoi(string(letter))
			if err == nil {
				return "", ErrInvalidString
			}
		}

		if currentLetter == 0 {
			count, err = strconv.Atoi(string(letter))
			if err != nil {
				currentLetter = letter
				continue
			}
			return "", ErrInvalidString
		}

		count, err = strconv.Atoi(string(letter))
		if err != nil {
			_, err := resultString.WriteRune(currentLetter)
			if err != nil {
				return "", ErrInvalidString
			}
			currentLetter = letter
			count = 0
			continue
		}

		_, err = resultString.WriteString(strings.Repeat(string(currentLetter), count))
		if err != nil {
			return "", ErrInvalidString
		}
		currentLetter = 0
	}

	_, err := resultString.WriteRune(currentLetter)
	if err != nil {
		return "", ErrInvalidString
	}

	return resultString.String(), nil
}
