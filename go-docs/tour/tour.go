// For keeping useful snippets from the tour of go tutorial.
package main

import "fmt"

func main() {
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
