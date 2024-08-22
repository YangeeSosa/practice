package kata

import (
	"sort"
	"strings"
	"unicode"
)

func Order(sentence string) string {
	if sentence == "" {
		return ""
	}
	words := strings.Fields(sentence)
	sort.Slice(words, func(i, j int) bool {
		return getDigit(words[i]) < getDigit(words[j])
	})

	return strings.Join(words, " ")
	return sentence
}

func getDigit(word string) int {
	for _, char := range word {
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
	}
	return 0
}
