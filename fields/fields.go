package fields

import (
	"fmt"
	"time"
)

type FieldType int

const (
	Int FieldType = iota
	String
	Char
	// TODO: Decimal
	Float
	Bool
	DateTime
)

type Field struct {
	tag       int
	val       interface{}
	fieldType FieldType
}

func (f *Field) GetType() FieldType {
	return f.fieldType
}

// Functions to create common Field types
func (f *Field) GetInt() (int, error) {
	if f.fieldType == Int {
		return f.val.(int), nil
	}
	return 0, fmt.Errorf("Field is of type %v not Int", f.fieldType)
}

func (f *Field) GetString() (string, error) {
	if f.fieldType == String {
		return f.val.(string), nil
	}
	return "", fmt.Errorf("Field is of type %v not String", f.fieldType)
}

func (f *Field) GetChar() (rune, error) {
	if f.fieldType == Char {
		return f.val.(rune), nil
	}
	return 0, fmt.Errorf("Field is of type %v not Char", f.fieldType)
}

func (f *Field) GetFloat() (float64, error) {
	if f.fieldType == Float {
		return f.val.(float64), nil
	}
	return 0, fmt.Errorf("Field is of type %v not Char", f.fieldType)
}

func (f *Field) GetBool() (bool, error) {
	if f.fieldType == Bool {
		return f.val.(bool), nil
	}
	return false, fmt.Errorf("Field is of type %v not Char", f.fieldType)
}

func (f *Field) GetDateTime() (time.Time, error) {
	if f.fieldType == DateTime {
		return f.val.(time.Time), nil
	}
	return time.Unix(0, 0), fmt.Errorf("Field is of type %v not DateTime", f.fieldType)
}

// Field accessors
func StringField(tag int, val string) *Field {
	return &Field{tag, val, String}
}

func IntField(tag int, val int) *Field {
	return &Field{tag, val, Int}
}

func CharField(tag int, val rune) *Field {
	return &Field{tag, val, Char}
}

func FloatField(tag int, val float64) *Field {
	return &Field{tag, val, Float}
}

func BoolField(tag int, val bool) *Field {
	return &Field{tag, val, Bool}
}

func DateTimeField(tag int, val time.Time) *Field {
	return &Field{tag, val, DateTime}
}

func (f *Field) String() string {
	return fmt.Sprintf("%v", f.val)
}
