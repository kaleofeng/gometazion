package ssh

import (
	"bytes"
	"sync"
)

type OutputBuffer struct {
	buffer bytes.Buffer
	mutex  sync.Mutex
}

func NewOutputBuffer() *OutputBuffer {
	return &OutputBuffer{}
}

func (s *OutputBuffer) Write(data []byte) (int, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.buffer.Write(data)
}

func (s *OutputBuffer) Shift() []byte {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	data := make([]byte, s.buffer.Len())
	copy(data, s.buffer.Bytes())
	s.buffer.Reset()
	return data
}
