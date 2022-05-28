package autocomplete

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/jmcveigh55/autocomplete/internal/dictionary"
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

	wordsDict := dictionary.NewDictionaryFromFile(fileName)
	wordsDict.Sort(wordsDict.SortWordsAscendStrategy)

	completionsDict := searchForCompletions(wordsDict, partialWord)
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

func searchForCompletions(dict *dictionary.Dictionary, partialWord string) *dictionary.Dictionary {
	completions := dictionary.NewDictionary()

	// linear search -> O(n)
	for _, entry := range dict.Entries {
		if strings.HasPrefix(entry.Word, partialWord) {
			completions.Entries = append(completions.Entries, entry)
		}
	}

	return completions
}
