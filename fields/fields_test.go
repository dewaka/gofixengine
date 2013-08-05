package fields

import (
	"testing"
)

func TestStringFields(t *testing.T) {
	s := StringField(1, "abc123")
	str, err := s.GetString()
	if err != nil {
		t.Errorf("Failed to retrieve string value")
	} else {
		if str != "abc123" {
			t.Errorf("Extracted value is not the one supplied")
		}
	}

	// Negative test - this should fail
	i, err2 := s.GetInt()
	if err2 == nil {
		t.Errorf("Field is of type String, but int value retrievable")
	} else {
		if i != 0 {
			t.Errorf("Did not get the default value 0 on failure")
		}
	}
}
