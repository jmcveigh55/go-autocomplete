package dictionary

import "testing"

func TestDictSortWordsAscendStrategy(t *testing.T) {
	dict := NewDictionary()
	dict.Entries = []*Entry{
		{Word: "elephant", Weight: 2},
		{Word: "cat", Weight: 1},
		{Word: "dog", Weight: 3},
	}

	dict.Sort(dict.SortWordsAscendStrategy)

	if dict.Entries[0].Word != "cat" {
		t.Errorf("Expected cat, got %s", dict.Entries[0].Word)
	}

	if dict.Entries[1].Word != "dog" {
		t.Errorf("Expected dog, got %s", dict.Entries[1].Word)
	}

	if dict.Entries[2].Word != "elephant" {
		t.Errorf("Expected elephant, got %s", dict.Entries[2].Word)
	}
}

func TestDictSortWordsDescendStrategy(t *testing.T) {
	dict := NewDictionary()
	dict.Entries = []*Entry{
		{Word: "elephant", Weight: 2},
		{Word: "cat", Weight: 1},
		{Word: "dog", Weight: 3},
	}

	dict.Sort(dict.SortWordsDescendStrategy)

	if dict.Entries[0].Word != "elephant" {
		t.Errorf("Expected elephant, got %s", dict.Entries[2].Word)
	}

	if dict.Entries[1].Word != "dog" {
		t.Errorf("Expected dog, got %s", dict.Entries[1].Word)
	}

	if dict.Entries[2].Word != "cat" {
		t.Errorf("Expected cat, got %s", dict.Entries[0].Word)
	}
}

func TestDictSortWeightsAscendStrategy(t *testing.T) {
	dict := NewDictionary()
	dict.Entries = []*Entry{
		{Word: "elephant", Weight: 2},
		{Word: "cat", Weight: 1},
		{Word: "dog", Weight: 3},
	}

	dict.Sort(dict.SortWeightsAscendStrategy)

	if dict.Entries[0].Word != "cat" {
		t.Errorf("Expected cat, got %s", dict.Entries[0].Word)
	}

	if dict.Entries[1].Word != "elephant" {
		t.Errorf("Expected elephant, got %s", dict.Entries[2].Word)
	}

	if dict.Entries[2].Word != "dog" {
		t.Errorf("Expected dog, got %s", dict.Entries[1].Word)
	}
}

func TestDictSortWeightsDescendStrategy(t *testing.T) {
	dict := NewDictionary()
	dict.Entries = []*Entry{
		{Word: "elephant", Weight: 2},
		{Word: "cat", Weight: 1},
		{Word: "dog", Weight: 3},
	}

	dict.Sort(dict.SortWeightsDescendStrategy)

	if dict.Entries[0].Word != "dog" {
		t.Errorf("Expected dog, got %s", dict.Entries[1].Word)
	}

	if dict.Entries[1].Word != "elephant" {
		t.Errorf("Expected elephant, got %s", dict.Entries[2].Word)
	}

	if dict.Entries[2].Word != "cat" {
		t.Errorf("Expected cat, got %s", dict.Entries[0].Word)
	}
}
