package fields

import (
	"fmt"
)

type FieldType int

const (
	Int FieldType = iota
	String
	Char
	Decimal
)

type Field struct {
	tag       int
	val       interface{}
	fieldType FieldType
}

func (f *Field) GetType() FieldType {
	return f.fieldType
}

func (f *Field) GetString() (string, error) {
	if f.fieldType == String {
		return f.val.(string), nil
	}
	return "", fmt.Errorf("Field is of type %v not String", f.fieldType)
}

func (f *Field) GetInt() (int, error) {
	if f.fieldType == Int {
		return f.val.(int), nil
	}
	return 0, fmt.Errorf("Field is of type %v not Int", f.fieldType)
}

func StringField(tag int, val string) *Field {
	return &Field{tag, val, String}
}

func IntField(tag int, val int) *Field {
	return &Field{tag, val, Int}
}
