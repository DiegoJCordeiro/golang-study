package main

import (
	"fmt"
)

var (
	length        int
	capacity      int
	sliceVariable = []string{"Go", "Language", "Study"}
)

func init() {
	length = len(sliceVariable)
	capacity = cap(sliceVariable)
}

func showInfoSlice() {

	fmt.Printf("Slice = %v\n", sliceVariable)
	fmt.Printf("Length = %d\n", length)
	fmt.Printf("Capacity = %d\n", capacity)
}

func manipulateSlice() {
	
	var removedFirstPosition = sliceVariable[1:]
	fmt.Printf("Removed first position = %v\n", removedFirstPosition)
	var removedLastPosition = sliceVariable[:capacity-1]
	fmt.Printf("Removed last position = %v\n", removedLastPosition)
	var appendedMoreOneData = append(sliceVariable, "Today")
	fmt.Printf("Appended more one data = %v\n", appendedMoreOneData)
}

func lesson4() {

	fmt.Println("Lesson 4 - Slices")
	showInfoSlice()
	manipulateSlice()
	fmt.Println("---")
}
