package onion

import "fmt"

type StepFunc func(*Pipeline)

type Pipeline struct {
	steps  []StepFunc
	cursor int
	err    error
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (m *Pipeline) Reset() {
	m.steps = make([]StepFunc, 0)
	m.cursor = -1
	m.err = nil
}

func (m *Pipeline) Start() {
	m.Next()
}

func (m *Pipeline) Use(stepFunc StepFunc) {
	m.steps = append(m.steps, stepFunc)
}

func (m *Pipeline) Next() {
	m.cursor++
	for m.cursor < len(m.steps) {
		m.steps[m.cursor](m)
		m.cursor++
	}
}

func (m *Pipeline) Abort(err error) {
	m.cursor = 1 << 8
	m.err = err
}

func (m *Pipeline) Cursor() int {
	return m.cursor
}

func (m *Pipeline) Result() error {
	return m.err
}

func (m *Pipeline) String() string {
	return fmt.Sprintf("{Steps:%v Cursor:%v Err:%v}", len(m.steps), m.cursor, m.err)
}
