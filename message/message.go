package message

import (
	"github.com/dewaka/gofixengine/fields"
)

type FieldMap struct {
	fmap map[int]*fields.Field
}

type Message struct {
	FieldMap
}

func NewMessage(val interface{}) *FieldMap {
	return &FieldMap{nil}
}
