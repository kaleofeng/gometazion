package task

type LongTask interface {
	GetSeq() int64
	SetSeq(int64)
	GetStage() int
	GetCode() int
	GetErr() error
	GetData() string
	IsComplete() bool
	IsSuccess() bool
	IsExpired() bool
	Execute()
}
