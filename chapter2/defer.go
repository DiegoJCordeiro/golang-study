package main

import "fmt"

func Lesson3() {

	fmt.Println("Lesson 3 - Defer")

	fmt.Printf("First Line, has defer: %t and your position is: %d \n", false, 1)
	defer fmt.Printf("Second Line, has defer: %t and your position is: %d \n", true, 2)
	fmt.Printf("Third Line, has defer: %t and your position is: %d \n", false, 3)
}
