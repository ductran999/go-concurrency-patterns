package main

import (
	"log/slog"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	ready := false

	var wg sync.WaitGroup

	for i := range 3 {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			cond.L.Lock()

			for !ready {
				slog.Info("Worker waiting...\n", "worker_id", workerID)
				cond.Wait()
			}

			slog.Info("-> Worker start work!\n", "worker_id", workerID)

			cond.L.Unlock()
		}(i)
	}

	slog.Info("Main: Prepare data...")
	time.Sleep(2 * time.Second)

	cond.L.Lock()
	ready = true
	// Wake up all go routine with broadcast
	cond.Broadcast()
	cond.L.Unlock()

	slog.Info("Main: Send Broadcast signal!")

	// wait for workers
	wg.Wait()
	slog.Info("Main: End.")
}
