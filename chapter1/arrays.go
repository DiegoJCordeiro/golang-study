package main

import (
	"fmt"
)

var (
	intArray [3]int
)

func init() {
	fillArray()
}

func fillArray() {

	intArray[0] = 1
	intArray[1] = 2
	intArray[2] = 3
}

func iterateArray() {

	for index, value := range intArray {
		fmt.Printf("ArrayPosition[%d] = Value %d\n", index, value)
	}
}

// Lesson3() is an introduction about arrays.
func Lesson3() {

	fmt.Println("Lesson 3 - Arrays")
	iterateArray()
	fmt.Println("---")
}
