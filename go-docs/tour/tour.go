// For keeping useful snippets from the tour of go tutorial.
package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
	"time"
)

// control flow statements
//    if, else
//    for - for loop (and while loop)
//    switch - the usual swtich-case statement
//    goto - review
//    defer- executes line when parent function returns
//    panic - stops execution of function, calls deferreds, moving up the stack until crashing the program.
//    recover - recovers a panic, (perhaps like catch in js try-catch)

func main() {
	// looping()
	// conditionals()
	// switching()
	// deferring()
	panicking()

}

func looping() {
	sum := 0
	// go has only one looping construct, no while, do, etc.
	// the init and post-opp conditions are optional...
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	// ...so, to do a "while" loop, and
	for sum < 1000 {
		sum += sum
	}
	// an infinite loop is expressed as
	for {
		fmt.Println("I'll never die!")
	}
}

// conditionals
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// ahhh, so go's "if" statements may contain an expression
// this allows for a ternary like conditional statement

func pow(x, n, lim float64) float64 {
	// js: x**n < lim ? v : lim
	if v := math.Pow(x, n); v < lim {
		// v scoped this if statement
		return v
	} else {
		// this template literal has codes for types, g appears to be float 32 and 64, but not int? Not sure, no internet here.
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func guessPow(x float64) float64 {
	z := float64(1)

	for i := 0; i < 100; i++ {
		y := z
		// this algo is called Newton's method. I feel less bad about not understanding it.
		z -= (z*z - x) / (2 * z)
		if y == z {
			fmt.Println("Same, Attempts: ", i)
			return z
		}
		if math.Abs(y-z) < .00000001 {
			fmt.Println("Close, Attempts: ", i)
			return z
		}
	}
	return z
}

func conditionals() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(guessPow(123456))
	fmt.Println(math.Sqrt(123456))
}

// V. similar to JS, except the break statements are implicit, and
// switch value can can be more varied types it stops running after matching.
func switching() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		//freebsd, openbsd, plan9, windows...
		fmt.Printf("%s. \n", os)
	}
	// switch w/o a condition is a boolean, so each case is a boolean expression
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// Defer, executes when the function returns
func deferring() {
	defer fmt.Println("world")
	// weirdly, it moves backwards so "hey hi" runs before "world"
	defer fmt.Println("hey hi")

	fmt.Println("hello")
}

// Defer example.
// can type both args with one type annotation?
// src == source
// dst == destination
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close() // defer is evaluated WHEN IT IS EXECUTED. so note if src is modified again later in the fn

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	written, err = io.Copy(dst, src)
	return
}

// panic
func panicking() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

/* Output from panicking()
Calling g.
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking!
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
Recovered in f 4
Returned normally from f.
*/
