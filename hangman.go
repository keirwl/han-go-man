package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

const min_word_length = 5
const max_guesses = 10

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

func print_guess_word(guess_word string) {
	for _, char := range guess_word {
		fmt.Printf("%c ", char)
	}
	fmt.Printf("\n")
}

func replaceAtIndex(str string, r rune, index int) string {
	return str[:index] + string(r) + str[index+1:]
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	word := random_word(dict_words())
	length := len(word)
	guess_word := strings.Repeat("_", length)

	for guesses := max_guesses; guesses > 0; {
		var guess string
		fmt.Println("\n")
		print_guess_word(guess_word)
		fmt.Print("Enter guess (letter or whole word): ")
		fmt.Scanln(&guess)

		if len(guess) > 1 {
			if guess == word {
				fmt.Printf("Correct! The word was %s!\n", word)
				break
			} else {
				guesses--
				fmt.Printf("Incorrect! Guesses remaining: %d\n\n", guesses)
			}
		} else {
			if strings.Contains(word, guess) {
				for i, char := range word {
					if char == []rune(guess)[0] {
						guess_word = replaceAtIndex(guess_word, char, i)
					}
				}
				if guess_word == word {
					fmt.Printf("Success! %s completes the word %s!\n", guess, word)
					break
				} else {
					fmt.Printf("%s is in the word.\n", guess)
				}
			} else {
				guesses--
				fmt.Printf("%s is not in the word. Guesses ramaining: %d\n\n", guess, guesses)
			}
		}
	}
	if guess_word != word {
		fmt.Printf("The word was %s\n", word)
	}
}
