package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}

	taskCh := make(chan Task, len(tasks))
	errCh := make(chan error, len(tasks))
	stopCh := make(chan struct{})

	for _, task := range tasks {
		taskCh <- task
	}
	close(taskCh)

	var errorsCount int
	var mu sync.Mutex

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for task := range taskCh {
				select {
				case <-stopCh:
					return
				default:
					if err := task(); err != nil {
						errCh <- err
						mu.Lock()
						errorsCount++
						if errorsCount >= m {
							mu.Unlock()
							select {
							case <-stopCh:
							default:
								close(stopCh)
							}
							return
						}
						mu.Unlock()
					}
				}
			}
		}()
	}
	wg.Wait()
	close(errCh)

	if errorsCount > 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}
