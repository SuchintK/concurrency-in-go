package ch3

import "fmt"

func Channels() {
	// Ways to create channels
	print := func(dataStream chan interface{}, i int) {
		go func() {
			dataStream <- i
		}()
	}

	// 1. Declaration & assignment
	var dataStream chan interface{}
	dataStream = make(chan interface{})
	print(dataStream, 1)
	fmt.Println("Data: ", <-dataStream)

	// 2. Walrus assignment
	dataStream2 := make(chan interface{})
	print(dataStream2, 2)
	fmt.Println("Data: ", <-dataStream2)
}
