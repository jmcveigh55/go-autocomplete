package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type entry struct {
	Word   string
	Weight int
}

type Dictionary struct {
	Entries []*entry
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

		dictionary.Entries = append(dictionary.Entries, &entry{Word: entryWord, Weight: entryWeight})

	}

	return dictionary
}

func (d *Dictionary) Print() {
	for i, entry := range d.Entries {
		fmt.Println("\t", i, entry.Word, entry.Weight)
	}
}

func (d *Dictionary) Sort(strategy func(i, j int) bool) {
	sort.Slice(d.Entries, strategy)
}

func (d *Dictionary) SortWordsAscendStrategy(i, j int) bool {
	return d.Entries[i].Word < d.Entries[j].Word
}

func (d *Dictionary) SortWordsDescendStrategy(i, j int) bool {
	return d.Entries[i].Word > d.Entries[j].Word
}

func (d *Dictionary) SortWeightsAscendStrategy(i, j int) bool {
	return d.Entries[i].Weight < d.Entries[j].Weight
}

func (d *Dictionary) SortWeightsDescendStrategy(i, j int) bool {
	return d.Entries[i].Weight > d.Entries[j].Weight
}
