package main

import "fmt"

func Lesson14() {

	fmt.Printf("Lesson 14 - Conversion Types \n")

	numberInt := 13
	numberFloat := float32(numberInt)

	fmt.Printf("Converted type %T to %T. The result is: %.2f\n", numberInt, numberFloat, numberFloat)

	numberInt = int(numberFloat)
	fmt.Printf("Converted type %T to %T. The result is: %d\n", numberFloat, numberInt, numberInt)
}
