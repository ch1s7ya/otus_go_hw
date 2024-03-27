package hw03frequencyanalysis

import (
	"cmp"
	"slices"
	"strings"
	"unicode"
)

type WordFrequency struct {
	Word  string
	Count int
}

func Top10(text string) []string {
	words := strings.Fields(text)

	words = toLowerAndTrim(words)

	frequencyListOfWords := buildFrequencyListOfWords(words)
	sortedListOfWords := sortWords(frequencyListOfWords)
	top10FrequentWords := getTop10(sortedListOfWords)

	var frequentWords []string
	for _, word := range top10FrequentWords {
		frequentWords = append(frequentWords, word.Word)
	}

	return frequentWords
}

func buildFrequencyListOfWords(words []string) map[string]int {
	frequenceListOfWords := make(map[string]int)
	for _, word := range words {
		_, ok := frequenceListOfWords[word]
		if ok {
			frequenceListOfWords[word]++
		} else {
			frequenceListOfWords[word] = 1
		}
	}
	return frequenceListOfWords
}

func getTop10(words []WordFrequency) []WordFrequency {
	if len(words) != 0 {
		return words[0:10]
	} else {
		return words
	}
}

func sortWords(words map[string]int) []WordFrequency {
	var sortedWords []WordFrequency

	for k, v := range words {
		sortedWords = append(sortedWords, WordFrequency{k, v})
	}

	slices.SortFunc(sortedWords, func(a, b WordFrequency) int {
		if n := cmp.Compare(b.Count, a.Count); n != 0 {
			return n
		}
		return cmp.Compare(a.Word, b.Word)
	})

	return sortedWords
}

func toLowerAndTrim(words []string) []string {
	var lowerCaseWords []string

	for _, word := range words {
		word = strings.ToLower(strings.TrimFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		}))

		if word != "" {
			lowerCaseWords = append(lowerCaseWords, word)
		}
	}

	return lowerCaseWords
}
