package hw05parallelexecution

import (
	"errors"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if len(tasks) == 0 || n < 0 {
		return nil
	}

	if m < 0 {
		return ErrErrorsLimitExceeded
	}

	tasksChan := make(chan Task, len(tasks))
	errChan := make(chan error, n)
	successChan := make(chan struct{}, n)

	doneChan := make(chan struct{})

	for i := 0; i < n; i++ {
		go func() {
			for {
				select {
				case <-doneChan:
					return
				default:
				}

				task, ok := <-tasksChan
				if !ok {
					return
				}
				if err := task(); err != nil {
					errChan <- err
				} else {
					successChan <- struct{}{}
				}
			}
		}()
	}

	var errCount, successCount int
	go func() {
		for {
			select {
			case <-errChan:
				errCount++
				if errCount >= m || errCount+successCount == len(tasks) {
					close(doneChan)
					return
				}
			case <-successChan:
				successCount++
				if errCount+successCount == len(tasks) {
					close(doneChan)
					return
				}
			}
		}
	}()

	for _, t := range tasks {
		tasksChan <- t
	}
	close(tasksChan)

	<-doneChan

	if errCount >= m && m > 0 {
		return ErrErrorsLimitExceeded
	} else {
		return nil
	}

}
