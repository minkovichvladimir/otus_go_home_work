package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m < 0 {
		return ErrErrorsLimitExceeded
	}

	tasksChan := make(chan Task, len(tasks))
	errChan := make(chan error, n)
	doneChan := make(chan struct{})

	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-doneChan:
					// закрывается в горутине подсчитывающей кол-во ошибок
					return
				default:
				}

				task, ok := <-tasksChan
				if !ok {
					close(doneChan)
					return
				}
				if err := task(); err != nil {
					errChan <- err
				}
			}
		}()
	}

	var errCount int
	go func() {
		for {
			select {
			case <-errChan:
				errCount++
				if errCount >= m {
					close(doneChan)
					return
				}
			case <-doneChan:
				// закрывается в горутинах воркерах в случае если из канала с тасками все вычитали
				return
			}
		}
	}()

	for _, t := range tasks {
		tasksChan <- t
	}
	close(tasksChan)

	wg.Wait()

	if errCount >= m && m > 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}
