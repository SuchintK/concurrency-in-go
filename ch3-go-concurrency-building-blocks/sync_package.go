package ch3

import (
	"fmt"
	"sync"
)

func WaitGroups() {
	var wg sync.WaitGroup
	var i = 0
	wg.Add(5)
	for i < 5 {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
		i++
	}
	wg.Wait()
	fmt.Print("Printing is complete")
}

func Mutexs() {
	var mu sync.Mutex
	counter := 0

	increment := func(wg *sync.WaitGroup) {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		counter++
		fmt.Println("Incrementing counter to ", counter)
	}

	decrement := func(wg *sync.WaitGroup) {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		counter--
		fmt.Println("Decrementing counter to ", counter)
	}

	var wg sync.WaitGroup

	// incrementing
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	// decrement
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go decrement(&wg)
	}

	wg.Wait()
	fmt.Println("Printing done")
}

func Dos() {
	var counter int

	increment := func() {
		counter++
	}

	var once sync.Once
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			once.Do(increment)
		}()
	}
	wg.Wait()
	fmt.Printf("The value of counter: %v\n", counter)
}

func Pools() {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}
	pool.Get()
	instance := pool.Get()
	pool.Put(instance)
	pool.Get() // Previous instance will be reused
}
