package autocomplete

import (
	"testing"

	"github.com/jmcveigh55/autocomplete/internal/dictionary"
)

func TestValidate(t *testing.T) {
	if !validate("a") {
		t.Errorf("Expected true, got false")
	}

	if !validate("hello") {
		t.Errorf("Expected true, got false")
	}

	if validate("hello777") {
		t.Errorf("Expected false, got true")
	}

	if validate("777hello777") {
		t.Errorf("Expected false, got true")
	}
}

func TestSearchForCompletions(t *testing.T) {
	dict := dictionary.NewDictionary()
	dict.Entries = []*dictionary.Entry{
		{Word: "elephant", Weight: 2},
		{Word: "cat", Weight: 1},
		{Word: "dinosaur", Weight: 4},
		{Word: "dog", Weight: 3},
	}

	// Trying to complete "snail" or "snake," both not in the dictionary.
	completions := searchForCompletions(dict, "sna")
	if len(completions.Entries) != 0 {
		t.Errorf("Expected 0 completion(s), got %d", len(completions.Entries))
	}

	// Trying to complete "elephant."
	completions = searchForCompletions(dict, "ele")
	if len(completions.Entries) != 1 {
		t.Errorf("Expected 1 completion(s), got %d", len(completions.Entries))
	}
	if completions.Entries[0].Word != "elephant" {
		t.Errorf("Expected elephant, got %s", completions.Entries[0].Word)
	}

	// Trying to complete "dog" and "dinosaur."
	completions = searchForCompletions(dict, "d")
	if len(completions.Entries) != 2 {
		t.Errorf("Expected 2 completion(s), got %d", len(completions.Entries))
	}
	if completions.Entries[0].Word != "dinosaur" {
		t.Errorf("Expected dinosaur, got %s", completions.Entries[0].Word)
	}
	if completions.Entries[1].Word != "dog" {
		t.Errorf("Expected dog, got %s", completions.Entries[1].Word)
	}
}
