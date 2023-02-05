package main

import (
	"fmt"
	"sync"

	"r00tdk/learngo/fizzBuzz"
)

func main() {
	var w sync.WaitGroup
	c1 := make(chan string)
	for i := 0; i < 100; i++ {
		w.Add(1)
		go fizzBuzz.GetResult(i, &w, c1)
	}

	var counter int = 0

	go func() {
		for {
			select {
			case v := <-c1:
				counter++
				fmt.Printf("Value retured on channel: %v\n", v)
			}
		}
	}()

	fmt.Printf("main thread is waiting .... \n")
	w.Wait()
}
