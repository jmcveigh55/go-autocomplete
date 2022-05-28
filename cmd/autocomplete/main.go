package main

import (
	"os"

	"github.com/jmcveigh55/autocomplete/internal/autocomplete"
)

func main() {
	if len(os.Args) != 3 {
		autocomplete.PrintUsage()
		os.Exit(1)
	}

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		autocomplete.PrintUsage()
		os.Exit(0)
	}

	wordsFile := os.Args[1]
	partialWord := os.Args[2]
	autocomplete.Run(wordsFile, partialWord)
}
