package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

type WordWithCount struct {
	Word  string
	Count int
}

func (wwc *WordWithCount) IncreaseCount() {
	if wwc == nil {
		return
	}

	wwc.Count++
}

func Top10(bigText string) []string {
	words := strings.Fields(bigText)
	wordsWithCounts := make([]WordWithCount, 0)

	for _, word := range words {
		founded := false
		for i, v := range wordsWithCounts {
			if v.Word == word {
				wordsWithCounts[i].IncreaseCount()
				founded = true
				break
			}
		}

		if !founded {
			wordsWithCounts = append(wordsWithCounts, WordWithCount{word, 1})
		}
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

	fmt.Println(result)
	return result
}
