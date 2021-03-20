package collection

type BitState uint64

func (bs *BitState) Set(bit int) *BitState {
	*bs |= 1 << bit
	return bs
}

func (bs *BitState) Clear(bit int) *BitState {
	*bs &= ^(1 << bit)
	return bs
}

func (bs *BitState) Test(bit int) bool {
	return 0 != *bs&(1<<bit)
}

func (bs *BitState) Value() uint64 {
	return uint64(*bs)
}
