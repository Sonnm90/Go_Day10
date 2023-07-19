package main

import (
	"demo_day_10/broadcast"
	"fmt"
	"sync"
	"time"
)

type Buffer struct {
	data     []int
	capacity int
	//mutex    sync.Mutex
	notEmpty sync.Cond
	notFull  sync.Cond
}

func NewBuffer(capacity int) *Buffer {
	b := &Buffer{
		data:     make([]int, 0, capacity),
		capacity: capacity,
	}
	b.notEmpty = sync.Cond{}
	b.notFull = sync.Cond{}
	return b
}

func (b *Buffer) Put(value int) {
	//b.mutex.Lock()
	b.notFull.L.Lock()
	//defer b.mutex.Unlock()
	defer b.notFull.L.Unlock()
	for len(b.data) == b.capacity {
		fmt.Println("waiting----->len data: ", len(b.data))

		b.notFull.Wait()
	}

	b.data = append(b.data, value)
	b.notEmpty.Signal()
}

func (b *Buffer) Get() int {
	//b.mutex.Lock()
	b.notEmpty.L.Lock()
	//defer b.mutex.Unlock()
	defer b.notEmpty.L.Unlock()
	for len(b.data) == 0 {
		fmt.Println("waiting------>len data: ", len(b.data))
		b.notEmpty.Wait()
	}

	value := b.data[0]
	b.data = b.data[1:]
	b.notFull.Signal()

	return value
}

func Signal() {
	buffer := NewBuffer(3)
	//buffer.notEmpty = sync.NewCond(buffer.notEmpty.L)
	//buffer.notFull = sync.NewCond(buffer.notFull.L)
	var wg sync.WaitGroup
	wg.Add(2)

	// Producer goroutine
	go func() {
		defer wg.Done()

		for i := 1; i <= 10; i++ {
			//time.Sleep(1 * time.Second)
			buffer.Put(i)
			fmt.Println("Produced:", i)
		}
	}()

	// Consumer goroutine
	go func() {
		defer wg.Done()

		for i := 1; i <= 10; i++ {
			time.Sleep(time.Second * 3)
			value := buffer.Get()
			fmt.Println("Consumed:", value)
		}
	}()

	wg.Wait()
}

func main() {
	//Signal
	//Signal()

	//Broadcast
	broadcast.Demo()
}
