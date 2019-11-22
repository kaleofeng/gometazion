package kit

import "sync/atomic"

// IdGenerator generate ids of int32 type
type IdGenerator struct {
	id int32
}

// New returns a new object
func New(id int32) *IdGenerator {
	return &IdGenerator{
		id: id,
	}
}

// NextId returns a new id
func (ig *IdGenerator) NextId() int32 {
	return atomic.AddInt32(&ig.id, 1)
}

// CurrentId returns the current id
func (ig *IdGenerator) CurrentId() int32 {
	return ig.id
}
