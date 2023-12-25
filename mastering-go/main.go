package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// print
	fmt.Println("Hello world.")
	fmt.Println("2 + 2 = ", 2+2)

	// create empty file
	myFile, es := os.Create("hello.txt")
	if es != nil {
		log.Fatal(es)
	}
	log.Println(myFile)
	myFile.Close()
}
