package middleware

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"

	e "gitlab.ozon.dev/yuweebix/homework-1/internal/middleware/errors"
)

const (
	maxGoroutines = 1_000_000
)

var (
	numJobs = atomic.Int64{}
)

// job представляет работу, передаваемое в пул воркеров
type job struct {
	ctx  context.Context
	task func()
	cmd  []string
}

// WorkerPool представляет пул воркеров, который управляет выполнением работ
type WorkerPool struct {
	ctx              context.Context
	cancel           context.CancelFunc
	pool             *sync.Pool
	mu               sync.Mutex
	wg               sync.WaitGroup
	notificationChan chan string
	numWorkers       int
}

// NewWorkerPool создает новый пул воркеров с заданным количеством воркеров и контекстом
func NewWorkerPool(ctx context.Context, numWorkers int, notificationChan chan string) (*WorkerPool, error) {
	// слишком много воркеров
	if numWorkers > maxGoroutines {
		return nil, e.ErrGoroutinesNumExceeded
	}

	ctx, cancel := context.WithCancel(ctx)
	return &WorkerPool{
		ctx:              ctx,
		cancel:           cancel,
		pool:             &sync.Pool{},
		notificationChan: notificationChan,
		numWorkers:       numWorkers,
	}, nil
}

// worker представляет собой функцию для выполнения заданий воркерами
func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for {
		// вытаскиваем работу из пула
		// если пул пустой, то продолжаем цикл
		// выходим по сигналу
		wp.mu.Lock()
		job, ok := wp.pool.Get().(*job) // работа сама удалится из пула
		wp.mu.Unlock()
		if !ok {
			select {
			case <-wp.ctx.Done():
				if numJobs.Load() == 0 {
					return
				}
			default:
			}
			continue
		}

		select {
		case <-job.ctx.Done():
			if numJobs.Load() == 0 {
				return
			}
		default:
		}

		isChecked := checkCmd(job.cmd)
		if isChecked {
			wp.notificationChan <- fmt.Sprintf("команда %v начала исполняться", job.cmd)
		}
		job.task()
		if isChecked {
			wp.notificationChan <- fmt.Sprintf("команда %v закончила исполняться", job.cmd)
		}

		numJobs.Add(-1)
	}
}

// Start запускает воркеров в пуле
func (wp *WorkerPool) Start() {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	for i := 1; i <= wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}
}

// Stop завершает всех воркеров в пуле
func (wp *WorkerPool) Stop() {
	wp.cancel()
	wp.wg.Wait()
	close(wp.notificationChan)
}

// Enqueue добавляет работу в пул
func (wp *WorkerPool) Enqueue(ctx context.Context, task func(), cmd []string) {
	numJobs.Add(1)

	job := &job{
		ctx:  ctx,
		task: task,
		cmd:  cmd,
	}

	wp.mu.Lock()
	wp.pool.Put(job)
	wp.mu.Unlock()
}

// AddWorkers добавляет новых воркеров в пул
func (wp *WorkerPool) AddWorkers(n int) error {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	// слишком много воркеров
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

// RemoveWorkers удаляет воркеров из пула
func (wp *WorkerPool) RemoveWorkers(n int) error {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	// не может быть меньше одного воркера
	if wp.numWorkers-n < 1 {
		return e.ErrGoroutinesNumSubceeded
	}

	for i := 0; i < n; i++ {
		// создаем контекст с отменой для завершения одного воркера
		jobCtx, cancel := context.WithCancel(wp.ctx)
		cancel()                           // вызываем cancel для завершения воркера
		wp.Enqueue(jobCtx, func() {}, nil) // отправляем dummy job с cancel для завершения воркера
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
