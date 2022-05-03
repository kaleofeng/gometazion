package id

import "sync/atomic"

// Generator generate id of int64 type.
type Generator struct {
	id int64
}

// NewGenerator new an instance.
func NewGenerator(id int64) *Generator {
	return &Generator{
		id: id,
	}
}

// CurrentId returns the current id.
func (s *Generator) CurrentId() int64 {
	return s.id
}

// NextId returns a new id.
func (s *Generator) NextId() int64 {
	return atomic.AddInt64(&s.id, 1)
}
