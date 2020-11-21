package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GameWord(wordChars [] string, inputChars [] string) int {
	ScreenText(1, 0, "")

	hitCount := 0
	for _, wordValue := range wordChars {
		wordCharSymbol, wordCharIsHit := showCharOrMask(wordValue, inputChars)
		fmt.Printf("%s ", wordCharSymbol)

		if wordCharIsHit {
			hitCount++
		}
	}

	ScreenText(0, 1, "")

	return hitCount
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

func WordInput() string {
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

