package main

import (
	"fmt"
	"runtime"
	"sort"
)

// functions in go are first class values, ref: https://dev.to/andyhaskell/basics-of-first-class-functions-in-go-d5p

// function taking taking functions as parameters
func slushie(flavours ...func()) {
	for _, flavour := range flavours {
		flavour()
	}
}

// function returning function
func slushieSpeedDial(flavours ...func()) func() {
	return func() {
		slushie(flavours...)
	}
}

func tickets(ticket int) func() int {
	// closure
	return func() int {
		pc, _, _, _ := runtime.Caller(0)
		ticket += 1
		fmt.Printf("ticket: %d, ticketAddr: %p, pc: 0x%x\n",
			ticket, &ticket, pc)
		return ticket
	}
}

func checkFuncs() {
	chocolate := func() { fmt.Println("Flavour is chocolate") }
	vanila := func() { fmt.Println("Flavour is vanila") }

	slushie(chocolate, vanila)
	blueberry := func() { fmt.Println("Flavour is blueberry") }
	strawberry := func() { fmt.Println("Flavour is strawberry") }
	blueAndStraw := slushieSpeedDial(blueberry, strawberry)

	blueAndStraw()

	/* sort.Slice() method also takes function as input
	func Slice(
		x interface{}, // slice we want to sort
		less func(i, j) bool // custom logic (return true if x[i] < x[j])
	) { // sorting logic }
	*/

	values := []string{"rest", "check", "this", "matter", "validate", "number"}
	// sort.Slice(values, func(i, j int) bool {
	// 	return values[i] < values[j]
	// })
	sort.Slice(values, func(i, j int) bool {
		return len(values[i]) < len(values[j])
	})
	fmt.Println(values)

	tickets1 := tickets(100)
	tickets2 := tickets(200)
	fmt.Printf("tickets1: %p, tickets2: %p\n", tickets1, tickets2)
	tickets1()
	tickets2()

}
