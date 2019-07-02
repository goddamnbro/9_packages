package homework

import (
	"log"
	"regexp"
	"sort"
	"strings"
)

func GetFrequencyWords(text string) []string {
	// convert text to slice
	words := strings.Split(text, " ")

	wordMap := make(map[string]int)

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	// create map with words, {"word": count}
	for _, word := range words {
		w := reg.ReplaceAllString(word, "")
		w = strings.ToLower(w)

		value, ok := wordMap[w]
		if !ok {
			wordMap[w] = 1
		} else {
			wordMap[w] = value + 1
		}
	}

	type kv struct {
		Key   string
		Value int
	}

	var wordSlice []kv
	for k, v := range wordMap {
		wordSlice = append(wordSlice, kv{k, v})
	}

	// we want to have identical result with the words with the same count
	sort.Slice(wordSlice, func(i, j int) bool {
		if wordSlice[i].Value == wordSlice[j].Value {
			return wordSlice[i].Key > wordSlice[j].Key
		}
		return wordSlice[i].Value > wordSlice[j].Value
	})

	var top10words []string
	// if our text don't have 10 words - our code will not fail :)
	// the second option - checking length of text
	length := 10
	if len(wordMap) < 10 {
		length = len(wordMap)
	}
	for i:=0; i < length; i++ {
		top10words = append(top10words, wordSlice[i].Key)
	}

	return top10words
}
