package broadcast

import (
	"fmt"
	"sync"
	"time"
)

func Demo() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	// Số lượng goroutine chờ
	numGoroutines := 5

	for i := 1; i <= numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			fmt.Printf("Goroutine %d: Waiting\n", id)
			mu.Lock()
			cond.Wait()
			mu.Unlock()

			fmt.Printf("Goroutine %d: Resumed\n", id)
		}(i)
	}

	// Tạo một goroutine để gửi broadcast sau 3 giây
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Sending broadcast signal")
		mu.Lock()
		cond.Broadcast()
		//cond.Signal()
		//cond.Signal()
		//cond.Signal()
		//cond.Signal()
		mu.Unlock()
	}()

	wg.Wait()
	fmt.Println("All goroutines resumed")
}
