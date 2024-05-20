package main

import (
	"fmt"
)

func changeWithoutUsingPointer(firstValue, secondValue int32) {

	firstValue = 20
	secondValue = 30
}

func changeUsingPointer(firstValue, secondValue *int32) {

	*firstValue = 20
	*secondValue = 30
}

func lesson8() {

	fmt.Println("Lesson 8 - Pointers")

	var firstValue int32 = 10
	var secondValue int32 = 10

	changeWithoutUsingPointer(firstValue, secondValue)
	fmt.Printf("Without Using Pointers - First Value = %d and SecondValue = %d \n", firstValue, secondValue)
	changeUsingPointer(&firstValue, &secondValue)
	fmt.Printf("Using Pointers - First Value = %d and SecondValue = %d \n", firstValue, secondValue)

	fmt.Println("---")
}
