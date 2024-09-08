package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		m = 0 // all tasks should be done well
	}

	at := &AtomicTasks{tasks: tasks}
	wg := &sync.WaitGroup{}

	ch := make(chan error, n)
	defer close(ch)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				if m == 0 && at.GetErrorCount() > m {
					ch <- ErrErrorsLimitExceeded
					return
				}

				if m != 0 && at.GetErrorCount() >= m {
					ch <- ErrErrorsLimitExceeded
					return
				}

				if err := at.Consume(); err != nil {
					ch <- err
					return
				}
			}
		}()
	}
	wg.Wait()

	for i := 0; i < n; i++ {
		if err := <-ch; err != nil {
			if !errors.Is(err, ErrTaskLimitExceeded) {
				return err
			}
		}
	}

	return nil
}
