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

func Closures1() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "world"
	}()
	wg.Wait()
	fmt.Print(salutation)
	// Value is modified
	// Conclusion: Goroutines execute in the same address space they were created in
}

func Closures2() {
	var wg sync.WaitGroup
	for _, w := range []string{"hello", "world", "!"} {
		wg.Add(1)
		go func(w string) {
			defer wg.Done()
			fmt.Print(w + " ")
		}(w)
	}
	wg.Wait()
}
