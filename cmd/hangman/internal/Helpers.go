package internal

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func HelperStringContains(list [] string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}

	return false
}

func HelperStringRandomElementArr(list [] string) string {
	rand.Seed(time.Now().Unix())

	return list[rand.Intn(len(list))]
}

func ScreenClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func ScreenText(topPadding int, bottomPadding int, message string) {
	for counter := 1; counter <= topPadding; counter++ {
		println()
	}

	println(message)

	for counter := 1; counter <= bottomPadding; counter++ {
		println()
	}
}

func WordInput(attempts int, maxAttempts int) string {
	ScreenText(0, 0, "")

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
