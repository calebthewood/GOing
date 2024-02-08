package main

import (
	"fmt"
	"log"
	"os"
)

var (
	myfile *os.FileInfo
	es     error
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
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

func partTwo() {
	// Stat() func returns the file in & if there is
	// no file, it returns error
	myfile, es := os.Stat("hello.txt")
	if es != nil {
		// checking if given file exists or not
		if os.IsNotExist(es) {
			log.Fatal("File not Found")
		}
	}
	log.Println("File Exists")
	log.Println("File Details:")
	log.Println("Name: ", myfile.Name())
	log.Println("Size: ", myfile.Size())
}

func partThree() {
	// this feels wrong, but chatGPT has assured me that it is
	// a normal way to handle errors in Go.
	if er := os.Mkdir("a", os.ModePerm); er != nil {
		log.Fatal(er)
	}
}
