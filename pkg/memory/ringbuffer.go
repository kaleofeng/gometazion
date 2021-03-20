package memory

// RingBuffer a ring circular buffer
type RingBuffer struct {
	buffer    []byte
	length    int
	pullIndex int
	pushIndex int
}

// NewRingBuffer returns a new ring buffer
func NewRingBuffer(capacity int) (pb *RingBuffer) {
	return &RingBuffer{
		make([]byte, capacity),
		0,
		0,
		0,
	}
}

// Push push data into the buffer
func (rb *RingBuffer) Push(inData []byte) int {
	inLength := len(inData)
	vacant := rb.GetPushLength()
	properLength := inLength
	if vacant < inLength {
		properLength = vacant
	}
	if properLength <= 0 {
		return 0
	}

	rightLength := rb.GetCapacity() - rb.pushIndex
	if properLength <= rightLength {
		copy(rb.buffer[rb.pushIndex:], inData[:properLength])
		rb.pushIndex += properLength
	} else {
		copy(rb.buffer[rb.pushIndex:], inData[:rightLength])
		copy(rb.buffer, rb.buffer[rightLength:properLength])
		rb.pushIndex = properLength - rightLength
	}

	rb.pushIndex %= rb.GetCapacity()
	rb.length += properLength
	return properLength
}

// Pull pull data from the buffer
func (rb *RingBuffer) Pull(outBuffer []byte) int {
	outLength := cap(outBuffer)
	engaged := rb.GetPullLength()
	properLength := outLength
	if engaged < outLength {
		properLength = engaged
	}
	if properLength <= 0 {
		return 0
	}

	rightLength := rb.GetCapacity() - rb.pullIndex
	if properLength <= rightLength {
		copy(outBuffer, rb.buffer[rb.pullIndex:rb.pullIndex+properLength])
		rb.pullIndex += properLength
	} else {
		copy(outBuffer, rb.buffer[rb.pullIndex:])
		copy(outBuffer[rightLength:], rb.buffer[:properLength-rightLength])
		rb.pullIndex = properLength - rightLength
	}

	rb.pullIndex %= rb.GetCapacity()
	rb.length -= properLength
	return properLength
}

// Peek peek data in the buffer
func (rb *RingBuffer) Peek(outBuffer []byte) int {
	outLength := cap(outBuffer)
	engaged := rb.GetPullLength()
	properLength := outLength
	if engaged < outLength {
		properLength = engaged
	}
	if properLength <= 0 {
		return 0
	}

	rightLength := rb.GetCapacity() - rb.pullIndex
	if properLength <= rightLength {
		copy(outBuffer, rb.buffer[rb.pullIndex:rb.pullIndex+properLength])
	} else {
		copy(outBuffer, rb.buffer[rb.pullIndex:])
		copy(outBuffer[rightLength:], rb.buffer[:properLength-rightLength])
	}

	return properLength
}

// Skip skip data of length in the buffer
func (rb *RingBuffer) Skip(outLength int) int {
	engaged := rb.GetPullLength()
	properLength := outLength
	if engaged < outLength {
		properLength = engaged
	}
	if properLength <= 0 {
		return 0
	}

	rb.pullIndex = (rb.pullIndex + properLength) % rb.GetCapacity()
	rb.length -= properLength
	return properLength
}

// IsEmpty check if the buffer is empty
func (rb *RingBuffer) IsEmpty() bool {
	return rb.length == 0
}

// IsFull check if the buffer is full
func (rb *RingBuffer) IsFull() bool {
	return rb.length == rb.GetCapacity()
}

// GetCapacity get the buffer capacity
func (rb *RingBuffer) GetCapacity() int {
	return cap(rb.buffer)
}

// GetPullLength get the buffer pull length
func (rb *RingBuffer) GetPullLength() int {
	return rb.length
}

// GetPushLength get the buffer push length
func (rb *RingBuffer) GetPushLength() int {
	return rb.GetCapacity() - rb.length
}
