package main

import (
	"fmt"
	"strings"
)

// HangmanGame structure
type HangmanGame struct {
	word       string   // Word to guess
	letters    []string // Correctly guessed letters
	tries      int      // Number of remaining tries
	maxTries   int      // Maximum number of tries
	hiddenWord string   // Representation of the hidden word with underscores
}

// NewGame creates a new Hangman game instance
func NewGame(word string, maxTries int) *HangmanGame {
	word = strings.ToUpper(word) // Convert the word to uppercase to make case insensitive
	hiddenWord := strings.Repeat("_", len(word))
	return &HangmanGame{
		word:       word,
		letters:    make([]string, 0),
		tries:      maxTries,
		maxTries:   maxTries,
		hiddenWord: hiddenWord,
	}
}

// GuessLetter processes a letter guess
func (j *HangmanGame) GuessLetter(letter string) {
	letter = strings.ToUpper(letter) // Convert letter to uppercase
	if len(letter) != 1 || contains(j.letters, letter) || !strings.ContainsAny(letter, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		fmt.Println("Invalid input. Please enter a single alphabet letter that hasn't been guessed yet.")
		return
	}

	if strings.Contains(j.word, letter) {
		j.letters = append(j.letters, letter)
		// Update the representation of the hidden word
		hiddenWord := ""
		for _, char := range j.word {
			if contains(j.letters, string(char)) {
				hiddenWord += string(char)
			} else {
				hiddenWord += "_"
			}
		}
		j.hiddenWord = hiddenWord
	} else {
		j.tries--
	}
}

// Won checks if the game has been won
func (j *HangmanGame) Won() bool {
	return j.hiddenWord == j.word
}

// Lost checks if the game has been lost
func (j *HangmanGame) Lost() bool {
	return j.tries <= 0
}

// Helper function to check if a slice contains a string
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Main function to run the Hangman game
func main() {
	game := NewGame("GOL", 5)

	for {
		fmt.Println("Remaining Tries:", game.tries)
		fmt.Println("Word:", game.hiddenWord)

		var letter string
		fmt.Print("Guess a letter: ")
		fmt.Scanln(&letter)

		game.GuessLetter(letter)

		if game.Won() {
			fmt.Println("Congratulations! You've won! The word was:", game.word)
			break
		} else if game.Lost() {
			fmt.Println("You've lost! The word was:", game.word)
			break
		}
	}
}
