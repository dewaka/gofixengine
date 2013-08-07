package fields

import (
	"fmt"
	"testing"
	"time"
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

func TestCharField(t *testing.T) {
	cc := 'C'
	s := CharField(33, cc)
	c, err := s.GetChar()
	if err != nil {
		t.Errorf("Failed to retrieve char value")
	} else {
		if c != cc {
			t.Errorf("Extracted value %c but instead got %c", cc, c)
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

func TestFloatField(t *testing.T) {
	ff := 3.1234
	s := FloatField(33, ff)
	f, err := s.GetFloat()
	if err != nil {
		t.Errorf("Failed to retrieve float64 value")
	} else {
		if f != ff {
			t.Errorf("Extracted value %f but instead got %f", ff, f)
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

func TestDateTimeField(t *testing.T) {
	dd := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	f := DateTimeField(68, dd)

	d, err := f.GetDateTime()
	if err != nil {
		t.Errorf("Failed to retrieve DateTime value")
	} else {
		if d != dd {
			t.Errorf("Extracted value %v but instead got %v", dd, d)
		}
	}

	// Negative test - this should fail
	str, err2 := f.GetString()
	if err2 == nil {
		t.Errorf("Field is of type Int, but string value retrievable")
	} else {
		if str != "" {
			t.Errorf("Did not get the default value \"\" on failure, instead got %s", str)
		}
	}
}

func TestBooleanField(t *testing.T) {
	bb := true
	s := BoolField(33, bb)
	b, err := s.GetBool()
	if err != nil {
		t.Errorf("Failed to retrieve boolean value")
	} else {
		if b != bb {
			t.Errorf("Extracted value %b but instead got %b", bb, b)
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

// Field to string should return the same value as Sprintf with "%v"
func TestFieldToString(t *testing.T) {
	f := FloatField(34, 3.41234)
	sf := fmt.Sprintf("%s", f)
	if str := fmt.Sprintf("%v", 3.41234); sf != str {
		t.Errorf("Expected %s, but we got %s", str, sf)
	}

	f2 := CharField(38, 'C')
	sf2 := fmt.Sprintf("%s", f2)
	if str := fmt.Sprintf("%v", 'C'); sf2 != str {
		t.Errorf("Expected %s, but we got %s", str, sf2)
	}
}
