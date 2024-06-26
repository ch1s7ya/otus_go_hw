package hw03frequencyanalysis

import (
	"cmp"
	"regexp"
	"slices"
	"strings"
)

type WordFrequency struct {
	Word  string
	Count int
}

var (
	reLeft  = regexp.MustCompile(`[[:punct:]]+$`)
	reRight = regexp.MustCompile(`^[[:punct:]]+`)
)

func Top10(text string) []string {
	words := strings.Fields(text)

	words = toLowerAndTrim(words)

	frequencyListOfWords := buildFrequencyListOfWords(words)
	sortedListOfWords := sortWords(frequencyListOfWords)
	top10FrequentWords := getTop10(sortedListOfWords)

	frequentWords := make([]string, 0, 10)
	for _, word := range top10FrequentWords {
		frequentWords = append(frequentWords, word.Word)
	}

	return frequentWords
}

func buildFrequencyListOfWords(words []string) map[string]int {
	frequenceListOfWords := make(map[string]int)
	for _, word := range words {
		frequenceListOfWords[word]++
	}
	return frequenceListOfWords
}

func getTop10(words []WordFrequency) []WordFrequency {
	if len(words) >= 10 {
		return words[0:10]
	}

	return words
}

func sortWords(words map[string]int) []WordFrequency {
	sortedWords := make([]WordFrequency, 0)

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
		if len(word) < 2 {
			continue
		}
		pureWord := strings.ToLower(reRight.ReplaceAllString(reLeft.ReplaceAllString(word, ""), ""))
		if len(pureWord) == 0 {
			lowerCaseWords = append(lowerCaseWords, word)
		} else {
			lowerCaseWords = append(lowerCaseWords, pureWord)
		}
	}

	return lowerCaseWords
}
