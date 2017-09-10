package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

const min_word_length = 5

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func dict_words() []string {
	file, err := ioutil.ReadFile("/usr/share/dict/words")
	check(err)
	words := strings.Split(string(file), "\n")
	return words
}

func random_word(words []string) string {
	length := len(words)
	var word string
	for len(word) < min_word_length {
		index := rand.Intn(length)
		word = words[index]
	}
	return word
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	word := random_word(dict_words())
	fmt.Println(word)
}
