package utils

import "strings"

func ToTitleCase(input string) string {
	words := strings.Split(input, " ")
	smallwords := " a an on the to in "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") && word != string(word[0]) {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}
