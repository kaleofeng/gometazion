package kit

import "sync/atomic"

// IdGenerator generate id of int32 type.
type IdGenerator struct {
	id int32
}

// NewIdGenerator new an instance.
func NewIdGenerator(id int32) *IdGenerator {
	return &IdGenerator{
		id: id,
	}
}

// CurrentId returns the current id.
func (ig *IdGenerator) CurrentId() int32 {
	return ig.id
}

// NextId returns a new id.
func (ig *IdGenerator) NextId() int32 {
	return atomic.AddInt32(&ig.id, 1)
}
