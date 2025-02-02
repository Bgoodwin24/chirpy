package main

import "strings"

func handlerCleanText(content string) string {
	var bannedWords = map[string]bool{
		"kerfuffle": true,
		"sharbert":  true,
		"fornax":    true,
	}

	loweredWords := strings.ToLower(content)
	splitLowerWords := strings.Split(loweredWords, " ")
	splitBaseWords := strings.Split(content, " ")
	cleanedWords := make([]string, 0, len(splitLowerWords))

	for i, word := range splitLowerWords {
		if bannedWords[word] {
			cleanedWords = append(cleanedWords, "****")
		} else {
			cleanedWords = append(cleanedWords, splitBaseWords[i])
		}
	}
	return strings.Join(cleanedWords, " ")
}
