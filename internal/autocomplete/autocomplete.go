package autocomplete

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/jmcveigh55/autocomplete/internal/dictionary"
)

func Run(fileName, partialWord string) {
	if !validate(partialWord) {
		fmt.Println("invalid partial word: ", partialWord)
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
