package autocomplete

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type entry struct {
	word   string
	weight int
}

type Dictionary struct {
	entries []*entry
}

func NewDictionary() *Dictionary {
	return &Dictionary{}
}

func NewDictionaryFromFile(fileName string) *Dictionary {
	dictionary := NewDictionary()

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")

		rawEntry := strings.Split(line, " ")
		entryWord := rawEntry[0]
		entryWeight, err := strconv.Atoi(rawEntry[1])
		if err != nil {
			fmt.Println("Encountered error while parsing file: ", fileName)
			fmt.Println(err)
			os.Exit(1)
		}

		dictionary.entries = append(dictionary.entries, &entry{word: entryWord, weight: entryWeight})

	}

	return dictionary
}

func (d *Dictionary) Print() {
	for i, entry := range d.entries {
		fmt.Println("\t", i, entry.word, entry.weight)
	}
}

func (d *Dictionary) Sort(strategy func(i, j int) bool) {
	sort.Slice(d.entries, strategy)
}

func (d *Dictionary) SortWordsAscendStrategy(i, j int) bool {
	return d.entries[i].word < d.entries[j].word
}

func (d *Dictionary) SortWordsDescendStrategy(i, j int) bool {
	return d.entries[i].word > d.entries[j].word
}

func (d *Dictionary) SortWeightsAscendStrategy(i, j int) bool {
	return d.entries[i].weight < d.entries[j].weight
}

func (d *Dictionary) SortWeightsDescendStrategy(i, j int) bool {
	return d.entries[i].weight > d.entries[j].weight
}
