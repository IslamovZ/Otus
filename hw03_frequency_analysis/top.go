package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordWithCount struct {
	Word  string
	Count int
}

func Top10(inputText string) []string {
	words := strings.Fields(inputText)
	wordsCounts := make(map[string]int)

	for _, word := range words {
		count, ok := wordsCounts[word]

		if ok {
			wordsCounts[word] = count + 1
		} else {
			wordsCounts[word] = 1
		}
	}

	wordsWithCounts := make([]WordWithCount, 0)

	for key, count := range wordsCounts {
		wordsWithCounts = append(wordsWithCounts, WordWithCount{Word: key, Count: count})
	}

	sort.Slice(wordsWithCounts, func(i, j int) bool {
		first := wordsWithCounts[i]
		second := wordsWithCounts[j]

		if first.Count == second.Count {
			return first.Word < second.Word
		}

		return first.Count > second.Count
	})

	result := []string{}
	lenArr := len(wordsWithCounts)
	for i := 0; i < 10 && i < lenArr; i++ {
		result = append(result, wordsWithCounts[i].Word)
	}

	return result
}
