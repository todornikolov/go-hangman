package main

import (
	"github.com/todornikolov/gogame/internal"
)

func main() {
	internal.ScreenClear()

	internal.ScreenText(0, 1, "Welcome to Hangman demo game written in Go, version 0.1")
	internal.ScreenText(0, 1,
		"Hangman is a paper and pencil guessing game for two or more players.\n" +
		"One player thinks of a word, phrase or sentence and the other(s) tries to guess\n" +
		"it by suggesting letters within a certain number of guesses.")

	internal.GameLoop(internal.WordFromDictionary("assets/dictionary.txt"),8)
}
