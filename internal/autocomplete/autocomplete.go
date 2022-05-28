package autocomplete

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func PrintUsage() {
	fmt.Println("Usage: autocomplete <word-file> <partial-word>")
}

func Run(fileName, partialWord string) {
	if !validate(partialWord) {
		PrintUsage()
		fmt.Println("Invalid partial word: ", partialWord)
		os.Exit(1)
	}

	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		PrintUsage()
		fmt.Println(err)
		os.Exit(1)
	}

	dictionary := NewDictionaryFromFile(fileName)
	dictionary.Sort(dictionary.SortWordsAscendStrategy)

	completionsDict := searchForCompletions(dictionary, partialWord)
	completionsDict.Sort(completionsDict.SortWeightsDescendStrategy)

	completionsDict.Print()
}

func validate(partialWord string) bool {
	for _, c := range partialWord {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

func searchForCompletions(dictionary *Dictionary, partialWord string) *Dictionary {
	completions := NewDictionary()

	for _, entry := range dictionary.entries {
		if strings.HasPrefix(entry.word, partialWord) {
			completions.entries = append(completions.entries, entry)
		}
	}

	return completions
}
