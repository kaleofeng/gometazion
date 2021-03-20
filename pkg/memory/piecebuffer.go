package memory

// PieceBuffer a piece of continuous memory
type PieceBuffer struct {
	buffer    []byte
	pullIndex int
	pushIndex int
}

// NewPieceBuffer returns a new piece buffer
func NewPieceBuffer(capacity int) (pb *PieceBuffer) {
	return &PieceBuffer{
		make([]byte, capacity),
		0,
		0,
	}
}

// Push push data into the buffer
func (pb *PieceBuffer) Push(inData []byte) int {
	pushLength := pb.GetPushLength()
	if pushLength <= 0 {
		return 0
	}

	inLength := len(inData)
	if pushLength < inLength {
		inLength = pushLength
	}
	copy(pb.buffer[pb.pushIndex:], inData)
	pb.pushIndex += inLength
	return inLength
}

// Pull pull data from the buffer
func (pb *PieceBuffer) Pull(outBuffer []byte) int {
	pullLength := pb.GetPullLength()
	if pullLength <= 0 {
		return 0
	}

	outLength := cap(outBuffer)
	if pullLength < outLength {
		outLength = pullLength
	}
	copy(outBuffer, pb.buffer[pb.pullIndex:pb.pullIndex+outLength])
	pb.pullIndex += outLength
	return outLength
}

// Peek peek data in the buffer
func (pb *PieceBuffer) Peek(outBuffer []byte) int {
	pullLength := pb.GetPullLength()
	if pullLength <= 0 {
		return 0
	}

	outLength := cap(outBuffer)
	if pullLength < outLength {
		outLength = pullLength
	}
	copy(outBuffer, pb.buffer[pb.pullIndex:pb.pullIndex+outLength])
	return outLength
}

// Skip skip data of length in the buffer
func (pb *PieceBuffer) Skip(length int) int {
	pullLength := pb.GetPullLength()
	if pullLength <= 0 {
		return 0
	}

	if pullLength < length {
		length = pullLength
	}
	pb.pullIndex += length
	return length
}

// Compact compact the buffer
func (pb *PieceBuffer) Compact() {
	if pb.pullIndex == 0 {
		return
	}

	if pb.pushIndex > pb.pullIndex {
		copy(pb.buffer, pb.buffer[pb.pullIndex:pb.pushIndex])
		pb.pushIndex -= pb.pullIndex
		pb.pullIndex = 0
	} else {
		pb.pushIndex = 0
		pb.pullIndex = 0
	}
}

// GetBuffer get the buffer data
func (pb *PieceBuffer) GetBuffer() []byte {
	return pb.buffer
}

// GetCapacity get the buffer capacity
func (pb *PieceBuffer) GetCapacity() int {
	return cap(pb.buffer)
}

// GetPullLength get the buffer pull length
func (pb *PieceBuffer) GetPullLength() int {
	return pb.pushIndex - pb.pullIndex
}

// GetPushLength get the buffer push length
func (pb *PieceBuffer) GetPushLength() int {
	return pb.GetCapacity() - pb.pushIndex
}
