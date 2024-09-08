package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrTaskLimitExceeded = errors.New("task limit exceeded")

type AtomicTasks struct {
	tasks       []Task
	mux         sync.Mutex
	errorsCount int
}

func (at *AtomicTasks) Consume() error {
	task, err := at.GetTask()
	if err != nil {
		return err // no more available tasks
	}

	if err = task(); err != nil {
		at.IncErrorCount()
	}

	return nil
}

func (at *AtomicTasks) GetTask() (Task, error) {
	at.mux.Lock()
	defer at.mux.Unlock()

	if len(at.tasks) == 0 {
		return nil, ErrTaskLimitExceeded
	}

	result := at.tasks[0]
	at.tasks = at.tasks[1:]

	return result, nil
}

func (at *AtomicTasks) IncErrorCount() {
	at.mux.Lock()
	defer at.mux.Unlock()
	at.errorsCount++
}

func (at *AtomicTasks) GetErrorCount() int {
	at.mux.Lock()
	defer at.mux.Unlock()
	return at.errorsCount
}
