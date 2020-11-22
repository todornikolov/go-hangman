package internal

import (
	"fmt"
)

func GameWord(wordChars [] string, inputChars [] string, attempts int) (int, string) {
	hitCount := 0
	displayWord := ""
	for _, wordValue := range wordChars {
		wordCharSymbol, wordCharIsHit := showCharOrMask(wordValue, inputChars)
		displayWord = fmt.Sprintf("%s %s", displayWord, wordCharSymbol)

		if wordCharIsHit {
			hitCount++
		}
	}

	return hitCount, displayWord
}

func GameScreen(attempts int, displayWord string) {
	ScreenClear()
	ScreenText(0, 1, "Welcome to Hangman demo game written in Go")
	ScreenText(0, 2,
		"Hangman is a paper and pencil guessing game for two or more players.\n" +
			"One player thinks of a word, phrase or sentence and the other(s) tries to guess\n" +
			"it by suggesting letters within a certain number of guesses.")
	ScreenText(0, 1, drawGraphics(attempts))
	ScreenText(0, 1, displayWord)
}

func showCharOrMask(value string, inputChars [] string) (string, bool) {
	for _, inputValue := range inputChars {
		if value == inputValue {
			return inputValue, true
		}
	}

	return "_", false
}

func GameState(wordChars [] string, inputChars [] string) bool {
	for _, wordValue := range wordChars {
		if ! HelperStringContains(inputChars, wordValue) {
			return false
		}
	}

	return true
}

func countAttempts(hits int, state int, attempts int) (int, int) {
	if hits == state {
		attempts++

		return state, attempts
	}

	return hits, attempts
}
