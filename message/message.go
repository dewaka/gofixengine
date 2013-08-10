package message

import (
	"github.com/dewaka/gofixengine/fields"
)

// FieldMap is the basic Field contianer used by Messages and Groups
// Should preserve field ordering
// Should have fast lookup based on tag value for fields
type FieldMap struct {
	fmap map[int]*fields.Field
}

type Message struct {
	FieldMap
}

func NewMessage(val interface{}) *FieldMap {
	return &FieldMap{nil}
}
