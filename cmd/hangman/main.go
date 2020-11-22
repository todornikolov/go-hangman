package main

import (
	"github.com/todornikolov/go-hangman/internal"
)

func main() {
	internal.GameLoop(internal.WordFromDictionary("assets/dictionary.txt"),8)
}
