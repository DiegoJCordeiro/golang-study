package main

import (
	"fmt"
)

func usingSwitch(first int) {
	switch first {
	case 1:
		fmt.Printf("Conditional (Switch): You're selecting case 1.\n")
	case 2:
		fmt.Printf("Conditional (Switch): You're selecting case 1.\n")
	default:
		fmt.Printf("Conditional (Switch): You're selecting case default.\n")
	}
}

func usingIf(first, second int) {
	if first < second {
		panic("Conditional (If): The second parameter is major than first")
	} else {
		fmt.Printf("Conditional (If): The parameters are OK.\n")
	}
}

// Lesson13 it`s about conditional structure.
func Lesson13() {

	fmt.Printf("Lesson 13 - Conditional Structure \n")
	usingIf(10, 9)
	usingSwitch(1)
	fmt.Printf("---\n")
}
