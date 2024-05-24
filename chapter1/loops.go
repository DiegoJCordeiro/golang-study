package main

import "fmt"

var (
	array = make([]int, 10)
)

func init() {
	// Default Loop For
	for index := 0; index < len(array); index++ {
		array[index] = index + 1
	}
}

func executeLooping() {

	fmt.Printf("For Each\n")
	// Loop For Each
	for _, value := range array {
		fmt.Printf("Value Iterated=%d\n", value)
	}
	fmt.Printf("---\n")

	var count int = 0

	fmt.Printf("Infinite For\n")
	// Loop Infinite For
	for {
		count++

		if count == len(array) {
			break
		}

		fmt.Printf("Value Iterated=%d\n", count)
	}
}

func Lesson12() {
	fmt.Printf("Lesson 12 - Loops \n")
	executeLooping()
	fmt.Printf("---\n")
}
