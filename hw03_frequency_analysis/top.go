package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
)

func Top10(source string) []string {
	words := normalizeWords(strings.Fields(source))
	frequencyWords := make(map[string]int)
	for _, word := range words {
		frequencyWords[word]++
	}

	sorted := sortByFrequency(frequencyWords)
	if len(sorted) > 10 {
		return sorted[:10]
	}

	return sorted
}

func normalizeWords(words []string) (result []string) {
	for _, word := range words {
		if word == "-" {
			continue
		}

		if unicode.IsPunct([]rune(word[0:1])[0]) {
			word = word[1:]
		}
		if unicode.IsPunct([]rune(word[len(word)-1:])[0]) {
			word = word[:len(word)-1]
		}

		if word == "" {
			continue
		}

		result = append(result, strings.ToLower(word))
	}

	return result
}

func sortByFrequency(frequencyWords map[string]int) []string {
	words := make([]string, 0, len(frequencyWords))
	for k := range frequencyWords {
		words = append(words, k)
	}
	sort.Slice(words, func(i, j int) bool {
		if frequencyWords[words[i]] == frequencyWords[words[j]] {
			return words[i] < words[j]
		}
		return frequencyWords[words[i]] > frequencyWords[words[j]]
	})

	return words
}
