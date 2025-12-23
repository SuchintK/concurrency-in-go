package ch3

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sayHello() {
	defer wg.Done()
	fmt.Println("hello")
}

// Routines demonstrates three ways to start goroutines.
func Routines() {
	// Three ways to define goroutines
	wg.Add(3)

	// 1. Placing the go keyword before a function
	go sayHello()

	// 2. Anonymous Functions
	go func() {
		defer wg.Done()
		fmt.Println("hello")
	}()

	// 3. Assign Function to a variable
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}
	go sayHello()

	wg.Wait()
}
