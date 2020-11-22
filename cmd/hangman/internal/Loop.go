package internal

import (
	"fmt"
	"strings"
)

func GameLoop(gameWordString string, gameMaxAttempts int) {
	hitCount := 0
	attemptCount := 0
	gameWordArray := strings.Split(gameWordString, "")
	inputKeyboard := [] string {}

	for {
		roundHits, displayWord := GameWord(gameWordArray, inputKeyboard, attemptCount)
		hitCount, attemptCount = countAttempts(roundHits, hitCount, attemptCount)

		GameScreen(attemptCount, displayWord)

		inputKeyboard = append(inputKeyboard, WordInput(attemptCount, gameMaxAttempts))

		if GameState(gameWordArray, inputKeyboard) {
			_, displayWord = GameWord(gameWordArray, inputKeyboard, attemptCount)
			GameScreen(attemptCount, displayWord)
			ScreenText(1, 1, "Congrats! You won!")
			break
		}

		if attemptCount >= gameMaxAttempts {
			attemptCount++
			_, displayWord = GameWord(gameWordArray, inputKeyboard, attemptCount)
			GameScreen(attemptCount, displayWord)
			ScreenText(
				1,
				1,
				fmt.Sprintf("Game over! The word was \"%s\". Try again.", gameWordString),
			)
			break
		}
	}
}
