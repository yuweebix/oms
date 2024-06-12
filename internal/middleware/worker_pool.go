package middleware

import (
	"context"
	"fmt"
	"sync"
)

// job представляет задание, передаваемое в пул рабочих
type job struct {
	ctx    context.Context
	cancel context.CancelFunc
	task   func()
}

// WorkerPool представляет пул рабочих, который управляет выполнением заданий
type WorkerPool struct {
	jobs             chan job
	numWorkers       int
	wg               sync.WaitGroup
	mu               sync.Mutex
	notificationChan chan string
	cancel           context.CancelFunc
	ctx              context.Context
}

// NewWorkerPool создает новый пул рабочих с заданным количеством рабочих и контекстом
func NewWorkerPool(ctx context.Context, numWorkers int, notificationChan chan string) *WorkerPool {
	ctx, cancel := context.WithCancel(ctx)
	return &WorkerPool{
		jobs:             make(chan job, 100), // рандомное значение
		numWorkers:       numWorkers,
		notificationChan: notificationChan,
		cancel:           cancel,
		ctx:              ctx,
	}
}

// worker представляет собой функцию для выполнения заданий рабочими
func (wp *WorkerPool) worker(workerNumber int) {
	defer wp.wg.Done()
	for {
		select {
		case <-wp.ctx.Done():
			return
		case job := <-wp.jobs:
			select {
			case <-job.ctx.Done():
				return
			default:
				wp.notificationChan <- fmt.Sprintf("рабочий %v начал работу", workerNumber)
				job.task()
				wp.notificationChan <- fmt.Sprintf("рабочий %v закончил работу", workerNumber)
			}
		}
	}
}

// Start запускает рабочих в пуле
func (wp *WorkerPool) Start() {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	for i := 1; i <= wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// Stop завершает всех рабочих в пуле
func (wp *WorkerPool) Stop() {
	wp.cancel()
	close(wp.jobs)
	wp.wg.Wait()
}

// Enqueue добавляет задание в очередь
func (wp *WorkerPool) Enqueue(ctx context.Context, task func()) {
	job := job{
		ctx:  ctx,
		task: task,
	}
	wp.jobs <- job
}

// EnqueueWithCancel добавляет задание в очередь с отменой
func (wp *WorkerPool) EnqueueWithCancel(ctx context.Context, cancel context.CancelFunc, task func()) {
	job := job{
		ctx:    ctx,
		cancel: cancel,
		task:   task,
	}
	wp.jobs <- job
}

// AddWorkers добавляет новых рабочих в пул
func (wp *WorkerPool) AddWorkers(n int) {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	for i := 1; i <= n; i++ {
		wp.wg.Add(1)
		go wp.worker(wp.numWorkers + i)
	}
	wp.numWorkers += n
}

// RemoveWorkers удаляет рабочих из пула
func (wp *WorkerPool) RemoveWorkers(n int) {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	for i := 0; i < n; i++ {
		// создаем контекст с отменой для завершения одного рабочего
		jobCtx, cancel := context.WithCancel(wp.ctx)
		cancel()                                        // вызываем cancel для завершения рабочего
		wp.EnqueueWithCancel(jobCtx, cancel, func() {}) // отправляем dummy job с cancel для завершения рабочего
	}
}
