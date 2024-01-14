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
	// panicking()
	// pointing()
	// structing()
	// arrayed()
	// slicing()
	// moreSlicing()
	// nilSlice()
	makeSlice()
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

/*
Pointers

	a pointer holds the memory address of a value
	"*" operator denotes the pointer's underlying value
	"&" operator generates a pointer to its operand
*/
func pointing() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

/* Output
42
21
73
*/

/*
	 Structs
		a collection of fields
		accessed using . notation
		when using a pointer, (*p).X == p.X
			Why use a pointer?
*/
func structing() {
	type Vertex struct {
		X int
		Y int
	}
	structA := Vertex{3, 4}
	fmt.Println(Vertex{1, 2}.Y)
	fmt.Println(structA.X)

	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1} // name syntax
		v3 = Vertex{}     // i guess ints default to 0, do all ints do this?
		p  = &Vertex{1, 2}
	)
	fmt.Println(v1, p, v2, v3)
}

/*
	 Arrays
		go arrays set their size
		arrays have a fixed size, set when created
		'var a [10]int' where a is an array of 10 ints
*/
func arrayed() {
	var n [10]int // initialized as a array of ten 0's

	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

/*
	 Slices
		a dynamically sized, flexible view into the elements of an array.
		wut?
		created by setting an upper and lower bound on an array
		apparently much more commonly used than arrays in go
		a slice is not a copy of the array, but modidying the slice
		will modify the array it describes
		low is inclusive, high is exclusive
*/
func slicing() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // high of 5, low of 1
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// slice literal
	// c := [3]bool{true, true, false}
	// d := []bool{true, true, false}

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	structy := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(structy)
	arr := [10]string{"hi"}
	fmt.Println("arr", arr)
}

// This is how you would manually increase the size of a slice (or array)
// by creating a new one, double the size, and copying it.
func doubleTheSizeOfASlice(s []byte) []byte {
	t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
	for i := range s {
		t[i] = s[i]
	}
	s = t
	return t
}

func betterSliceSizeIncrease(s []byte) {
	t := make([]byte, len(s), (cap(s)+1)*2)
	copy(t, s)
	s = t
}

// taken from go docs
// whew, there is a proper append function
func AppendByte(slice []byte, data ...byte) []byte {
	sliceLength := len(slice)
	allLength := sliceLength + len(data)

	if allLength > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (allLength+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:allLength]
	copy(slice[sliceLength:allLength], data)
	return slice
}

// Filter returns a new slice holding only
// the elements of s that satisfy fn()
func Filter(slice []int, predicate func(int) bool) []int {
	var p []int // == nil
	for _, v := range slice {
		if predicate(v) {
			p = append(p, v)
		}
	}
	return p
}

func moreSlicing() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// grow it past it's capacity
	s = s[2:]
	s = append(s, 1, 2, 3, 4, 5)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func nilSlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func makeSlice() {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
