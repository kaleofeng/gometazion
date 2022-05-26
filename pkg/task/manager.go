package task

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type Manager struct {
	curSeq          int64
	mutex           sync.RWMutex
	taskMap         map[int64]LongTask
	taskChan        chan LongTask
	cancelCtx       context.Context
	cancelFunc      context.CancelFunc
	recycleInternal time.Duration
}

func NewManager() *Manager {
	return &Manager{
		curSeq:          1000,
		taskMap:         make(map[int64]LongTask),
		taskChan:        make(chan LongTask, 4),
		recycleInternal: 5 * time.Second,
	}
}

func (s *Manager) Init() error {
	ctx := context.Background()
	s.cancelCtx, s.cancelFunc = context.WithCancel(ctx)

	for i := 0; i < 4; i++ {
		go s.process(s.cancelCtx)
	}
	go s.recycle(s.cancelCtx)
	return nil
}

func (s *Manager) Get(seq int64) LongTask {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if task, ok := s.taskMap[seq]; ok {
		return task
	}
	return nil
}

func (s *Manager) Push(task LongTask) {
	task.SetSeq(s.getNextSeq())
	s.put(task)
	s.taskChan <- task
}

func (s *Manager) process(cancelCtx context.Context) {
	for {
		select {
		case <-cancelCtx.Done():
			return
		case task := <-s.taskChan:
			task.Execute()
		}
	}
}

func (s *Manager) recycle(cancelCtx context.Context) {
	tick := time.NewTicker(s.recycleInternal)
	defer tick.Stop()

	for {
		select {
		case <-cancelCtx.Done():
			return
		case <-tick.C:
			s.removeExpired()
		}
	}
}

func (s *Manager) getNextSeq() int64 {
	return atomic.AddInt64(&s.curSeq, 1)
}

func (s *Manager) put(task LongTask) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.taskMap[task.GetSeq()] = task
}

func (s *Manager) removeExpired() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for seq, task := range s.taskMap {
		if task.IsExpired() {
			delete(s.taskMap, seq)
		}
	}
}
