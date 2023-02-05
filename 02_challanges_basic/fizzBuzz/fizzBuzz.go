package fizzBuzz

import (
	"fmt"
	"sync"
)

func GetResult(a int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	if a%3 == 0 && a%5 == 0 {
		ch <- "FizzBuzz"
	} else if a%5 == 0 {
		ch <- "Buzz"
	} else if a%3 == 0 {
		ch <- "Fizz"
	} else {
		ch <- fmt.Sprintf("%d", a)
	}
}
