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
			t.Errorf("Did not get the default value 0 on failure, instead got %d", i)
		}
	}
}

func TestIntFields(t *testing.T) {
	num := 349
	s := IntField(900, num) // tag is 900, value is num
	n, err := s.GetInt()
	if err != nil {
		t.Errorf("Failed to retrieve int value")
	} else {
		if n != num {
			t.Errorf("Extracted value %d but instead got %d", num, n)
		}
	}

	// Negative test - this should fail
	str, err2 := s.GetString()
	if err2 == nil {
		t.Errorf("Field is of type Int, but string value retrievable")
	} else {
		if str != "" {
			t.Errorf("Did not get the default value \"\" on failure, instead got %s", str)
		}
	}
}
