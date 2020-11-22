package internal

import "fmt"

func drawGraphics(stage int) string {
	switch stage {
		case 2:
			return stageTwo()
		case 3:
			return stageThree()
		case 4:
			return stageFour()
		case 5:
			return stageFive()
		case 6:
			return stageSix()
		case 7:
			return stageSeven()
		case 8:
			return stageEight()
		case 9:
			return stageNine()
	}

	return stageDefault()
}

func stageDefault() string {
	return fmt.Sprint(
		" +---+\n" +
			" |\n" +
			" |\n" +
			" |\n" +
			" |\n" +
			" |\n",
	)
}

func stageTwo() string {
	return fmt.Sprint(
		" +---+\n" +
			" |   |\n" +
			" |\n" +
			" |\n" +
			" |\n" +
			" |\n",
	)
}

func stageThree() string {
	return fmt.Sprint(
		" +---+\n" +
			" |   |\n" +
			" | (*.*)\n" +
			" |\n" +
			" |\n" +
			" |\n",
	)
}

func stageFour() string {
	return fmt.Sprint(
		" +---+\n" +
			" |   |\n" +
			" | (*.*)\n" +
			" |   |\n" +
			" |\n" +
			" |\n",
	)
}

func stageFive() string {
	return fmt.Sprint(
		" +---+\n" +
			" |   |\n" +
			" | (*.*)\n" +
			" |  /|\n" +
			" |\n" +
			" |\n",
	)
}

func stageSix() string {
	return fmt.Sprint(
		" +---+\n" +
			" |   |\n" +
			" | (*.*)\n" +
			" |  /|\\\n" +
			" |\n" +
			" |\n",
	)
}

func stageSeven() string {
	return fmt.Sprint(
		" +---+\n" +
			" |   |\n" +
			" | (*.*)\n" +
			" |  /|\\\n" +
			" |  /\n" +
			" |\n",
	)
}

func stageEight() string {
	return fmt.Sprint(
		" +---+\n" +
			" |   |\n" +
			" | (*.*)\n" +
			" |  /|\\\n" +
			" |  / \\\n" +
			" |\n",
	)
}

func stageNine() string {
	return fmt.Sprint(
		" +---+\n" +
			" |   |\n" +
			" | (x_x)\n" +
			" |  /|\\\n" +
			" |  / \\\n" +
			" |\n",
	)
}
