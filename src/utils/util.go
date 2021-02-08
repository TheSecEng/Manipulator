package utils

import (
	"fmt"
	"strings"
)

func ToTitleCase(input string, rules string) string {
	formattedRules := strings.Join(strings.Split(rules, ","), " ")
	formattedRules = fmt.Sprintf(" %s ", formattedRules)
	words := strings.Split(input, " ")
	for index, word := range words {
		if strings.Contains(formattedRules, fmt.Sprintf(" %s ", word)) && word != string(word[0]) {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}
