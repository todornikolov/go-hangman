package internal

import (
	"log"
	"math/rand"
	"os"
	"os/exec"
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
