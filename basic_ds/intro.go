package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

type MyFloat float64

/*
Types of variables in Go:

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

func compute(fn func(float64, float64) float64) float64 {
	return 2 * fn(3, 4)
}

const Pi = 3.14 // := is not allowed

func basicStuff() {
	var i, j int = 1, 2
	fmt.Printf("i: %T/%v, j: %T/%v\n", i, i, j, j)
	var f float64 = math.Sqrt(float64(i*i + j*j))
	var z uint = uint(f)
	fmt.Printf("f: %v, z: %v\n", f, z)

	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}

	// init and post statements are optional
	for sum < 15 {
		sum += sum
	}

	/*
		for {
			// infinite loop
		}
	*/
	fmt.Println(sum)

	if v := math.Pow(2, 2); v < 5 { // v is scoped in if block
		fmt.Println(v)
	}

	if v := math.Pow(2, 3); v < 5 { // v is scoped till else block
		fmt.Println("<5: ", v)
	} else {
		fmt.Println(">5: ", v)
	}

	fmt.Printf("Go runs on ")
	// swtich without condition is switch true
	switch os := runtime.GOOS; os { // Need not to be constant
	case "darwin": // case condition need not to be constant
		fmt.Println("OS X.") // Implict break statement in Go
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s\n", os)
	}

	var k int = 10
	var ptr *int = &k
	*ptr = 12

	fmt.Println("value is ", *ptr)

	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}  // X: 1, Y:0(implicit)
		v3 = Vertex{}      // X: 0, Y:0
		p  = &Vertex{1, 2} // pointer to p
	)

	fmt.Println(v1, v2, v3, p)

	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	primes := [4]int{2, 3, 5, 7}
	var s []int = primes[0:3] // slice
	fmt.Println(primes, s)

	// Creates an array and then creates slice to reference it
	st := []struct {
		x, y int
	}{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}
	fmt.Println(st)

	sl := []int{1, 2, 3, 4, 5, 6}
	sl = sl[:0]
	fmt.Printf("len(sl): %v, cap(sl): %v\n", len(sl), cap(sl)) // cap is size of underlying array
	sl = sl[:4]
	fmt.Printf("len(sl): %v, cap(sl): %v\n", len(sl), cap(sl))
	sl = sl[2:]
	fmt.Printf("len(sl): %v, cap(sl): %v\n", len(sl), cap(sl))
	var sll []int // nil slice has value nil, len/cap both 0 with no underlying array
	fmt.Printf("len(sll): %v, cap(sll): %v\n", len(sll), cap(sll))

	sli := make([]int, 2, 5) // make (dynamically sized array) type, len, cap
	fmt.Printf("len(sli): %v, cap(sli): %v\n", len(sli), cap(sli))

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board = append(board, []string{"-", "-"})
	fmt.Println(board)

	for i, v := range sl {
		fmt.Printf("i: %v, v: %v\n", i, v)
	}

	m := make(map[string]Vertex)
	m["Test"] = Vertex{1, 2}
	fmt.Println(m, "with key ref:", m["Test"])
	ma := map[string]Vertex{
		"Check": Vertex{
			3, 4,
		},
	}
	fmt.Println(ma, "with key ref:", ma["Check"])
	mapp := map[string]Vertex{
		"Testing":  {5, 6},
		"Testingg": {7, 8},
	}
	fmt.Println(mapp, "with key ref:", mapp["Testing"])
	if v, ok := mapp["Testingg"]; ok {
		fmt.Println("Key is present with value: ", v)
	}
	delete(mapp, "Testingg")

	hypot := func(x, y float64) float64 {
		return x + y
	}
	fmt.Println(hypot(2.2, 3.3))
	fmt.Println("From Compute", compute(hypot))

	// Both pos and idx function closures point to different sum variables
	pos := addr()
	idx := addr()
	for i := 1; i < 5; i++ {
		fmt.Println("pos -> sum: ", pos(i))
		fmt.Println("idx -> sum: ", idx(i*2))
	}

	fn := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci series is: %d\n", fn())
	}

	v := Vertex{5, 5}
	fmt.Println("mul is ", v.Mul())

	var fv MyFloat = -3.2
	fvv := MyFloat(-5.10)
	fmt.Println("abs of myfloats are ", fv.Abs(), fvv.Abs())

	// For recursion in closures, need to declare variable with var first
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println("Fib(7): ", fib(7))

	var ab, bc Abser
	fll := MyFloat(-10.5)
	ab = fll
	fmt.Println("ab div is ", ab.Div())

	vert := Vertex{75, 15}
	bc = &vert
	fmt.Println("bc div is ", bc.Div())

	var in interface{}
	describe(in)

	in = 40
	describe(in)

	in = "Hello"
	describe(in)

	var intr interface{} = "Hello"

	// type assertion provides access to interface underlying concrete value
	val, ok := intr.(string)
	fmt.Println(val, ok)

	tt, ok := intr.(float64)
	fmt.Println(tt, ok)

	findType(21)
	findType("Hello")
	findType(true)

	person1 := Person{"Dheeraj", 29}
	person2 := Person{"Mukesh", 30}
	fmt.Println(person1, person2)

	mp := map[string]IPAddr{
		"LocalHost":  {127, 0, 0, 1},
		"Google DNS": {8, 8, 8, 8},
	}
	for _, v := range mp {
		fmt.Printf("%v \n", v)
	}

	if err := run(); err != nil {
		fmt.Printf("%s\n", err)
	}
}

// Closure is like having same function keeping memory(sum) in different scope for callers
func addr() func(int) int {
	sum := 0

	return func(x int) int {
		fmt.Println("Sum inside closure is ", sum)
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	first := 0
	second := 1

	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func (v Vertex) Mul() int {
	return v.X * v.Y
}

// We can declare a method on non-struct types, too.
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyFloat) Div() float64 {
	return float64(f * 11.0)
}

func (v *Vertex) Div() float64 {
	return float64(v.X / v.Y)
}

func describe(i interface{}) {
	fmt.Printf("(%[1]v, %[1]T)\n", i)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%[1]v is of type: %[1]T\n", v)
	case string:
		fmt.Printf("%[1]v is of type: %[1]T\n", v)
	default:
		fmt.Printf("%[1]v is of type: %[1]T\n", v)
	}
}

/*
	type Stinger interface {
		String() string
	}

>> fmt package looks for String() method implementation to print
*/
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func (ip IPAddr) String() string {
	return fmt.Sprintf("\"%v.%v.%v.%v\"", ip[0], ip[1], ip[2], ip[3])
}

/*
	type error interface {
		Error() string
	}
>> fmt package looks for Error() method implementation to print error
*/

func (e *MyError) Error() string {
	return fmt.Sprintf("Error: %v occured at %v", e.What, e.When)
}

func run() error {
	return &MyError{
		When: time.Now(),
		What: "Failure state",
	}
}
