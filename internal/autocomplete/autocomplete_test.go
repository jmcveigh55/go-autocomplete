package autocomplete

import "testing"

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
