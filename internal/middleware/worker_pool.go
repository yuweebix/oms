package middleware

import (
	"context"
	"fmt"
	"strings"
	"sync"

	e "gitlab.ozon.dev/yuweebix/homework-1/internal/middleware/errors"
)

const (
	maxGoroutines = 1_000_000
)

// job представляет задание, передаваемое в пул рабочих
type job struct {
	ctx    context.Context
	cancel context.CancelFunc
	task   func()
	cmd    []string
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
func NewWorkerPool(ctx context.Context, numWorkers int, notificationChan chan string) (*WorkerPool, error) {
	// слишком много рабочих
	if numWorkers > maxGoroutines {
		return nil, e.ErrGoroutinesNumExceeded
	}

	ctx, cancel := context.WithCancel(ctx)
	return &WorkerPool{
		jobs:             make(chan job, 100), // рандомное значение
		numWorkers:       numWorkers,
		notificationChan: notificationChan,
		cancel:           cancel,
		ctx:              ctx,
	}, nil
}

// worker представляет собой функцию для выполнения заданий рабочими
// args передаются снова, чтобы выписать, какая команда работает
func (wp *WorkerPool) worker() {
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
				if checkCmd(job.cmd) {
					wp.notificationChan <- fmt.Sprintf("команда %v начала исполняться", job.cmd)
				}
				job.task()
				if checkCmd(job.cmd) {
					wp.notificationChan <- fmt.Sprintf("команда %v закончила исполняться", job.cmd)
				}
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
		go wp.worker()
	}
}

// Stop завершает всех рабочих в пуле
func (wp *WorkerPool) Stop() {
	wp.cancel()
	wp.wg.Wait()
	close(wp.jobs)
	close(wp.notificationChan)
}

// Enqueue добавляет задание в очередь
func (wp *WorkerPool) Enqueue(ctx context.Context, task func(), cmd []string) {
	job := job{
		ctx:  ctx,
		task: task,
		cmd:  cmd,
	}
	wp.jobs <- job
}

// EnqueueWithCancel добавляет задание в очередь с отменой
func (wp *WorkerPool) EnqueueWithCancel(ctx context.Context, cancel context.CancelFunc, task func(), cmd []string) {
	job := job{
		ctx:    ctx,
		cancel: cancel,
		task:   task,
		cmd:    cmd,
	}
	wp.jobs <- job
}

// AddWorkers добавляет новых рабочих в пул
func (wp *WorkerPool) AddWorkers(n int) error {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	// слишком много рабочих
	if wp.numWorkers+n > maxGoroutines {
		return e.ErrGoroutinesNumExceeded
	}

	for i := 1; i <= n; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}
	wp.numWorkers += n

	return nil
}

// RemoveWorkers удаляет рабочих из пула
func (wp *WorkerPool) RemoveWorkers(n int) error {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	// не может быть меньше одного рабочего
	if wp.numWorkers-n < 1 {
		return e.ErrGoroutinesNumSubceeded
	}

	for i := 0; i < n; i++ {
		// создаем контекст с отменой для завершения одного рабочего
		jobCtx, cancel := context.WithCancel(wp.ctx)
		cancel()                                             // вызываем cancel для завершения рабочего
		wp.EnqueueWithCancel(jobCtx, cancel, func() {}, nil) // отправляем dummy job с cancel для завершения рабочего
	}
	wp.numWorkers -= n

	return nil
}

func checkCmd(cmd []string) bool {
	if len(cmd) == 0 {
		return false
	}
	if strings.Contains(strings.Join(cmd, " "), "help") {
		return false
	}
	if strings.Contains(strings.Join(cmd, " "), "-h") {
		return false
	}
	return true
}
