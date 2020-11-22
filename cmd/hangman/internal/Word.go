package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GameWord(wordChars [] string, inputChars [] string) int {
	GameScreen()

	hitCount := 0
	for _, wordValue := range wordChars {
		wordCharSymbol, wordCharIsHit := showCharOrMask(wordValue, inputChars)
		fmt.Printf("%s ", wordCharSymbol)

		if wordCharIsHit {
			hitCount++
		}
	}

	return hitCount
}

func GameScreen() {
	ScreenClear()
	ScreenText(0, 1, "Welcome to Hangman demo game written in Go, version 0.1")
	ScreenText(0, 2,
		"Hangman is a paper and pencil guessing game for two or more players.\n" +
			"One player thinks of a word, phrase or sentence and the other(s) tries to guess\n" +
			"it by suggesting letters within a certain number of guesses.")
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

func WordInput(attempts int, maxAttempts int) string {
	ScreenText(1, 0, "")

	fmt.Printf("(%d/%d) Try a letter: ", attempts, maxAttempts)

	reader := bufio.NewReader(os.Stdin)
	inputKeyboardTry, _ := reader.ReadString('\n')

	return strings.TrimSuffix(inputKeyboardTry, "\n")
}

func WordFromDictionary(dictionaryPath string) string {
	file, err := os.Open(dictionaryPath)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var lines [] string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return HelperStringRandomElementArr(lines)
}

