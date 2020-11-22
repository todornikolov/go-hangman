package internal

import (
	"fmt"
	"strings"
)

func GameLoop(gameWordString string, gameMaxAttempts int) {
	attemptCount := 1
	hitCount := 0
	gameWordArray := strings.Split(gameWordString, "")
	inputKeyboard := [] string {}

	GameWord(gameWordArray, inputKeyboard)

	for {
		inputKeyboard = append(inputKeyboard, WordInput(attemptCount, gameMaxAttempts))

		hitCount, attemptCount = countAttempts(GameWord(gameWordArray, inputKeyboard), hitCount, attemptCount)

		if GameState(gameWordArray, inputKeyboard) {
			ScreenText(2, 1, "Congrats! You won!")
			break
		}

		if attemptCount >= gameMaxAttempts {
			ScreenText(
				2,
				1,
				fmt.Sprintf("Game over! The word was \"%s\". Try again.", gameWordString),
			)

			break
		}
	}
}

func countAttempts(hits int, state int, attempts int) (int, int) {
	if hits == state {
		attempts++

		return state, attempts
	}

	return hits, attempts
}
