package main

import "fmt"

var (
	firstGeneric  string
	secondGeneric int32
)

type GenericImplementer struct {
	name string
}

type Generics interface{}

func processGeneric(generic Generics) {
	fmt.Printf("Processing generic %T: %+v\n", generic, generic)
}

// Function return Type Assertion
func convertGenericToInt(generic Generics) int32 {
	return generic.(int32)
}

// Function return Type Assertion
func convertGenericToString(generic Generics) string {
	return generic.(string)
}

func lesson10() {

	fmt.Println("Lesson 10 - Generics")
	firstGeneric = "Hello Generics"
	secondGeneric = 0
	thirdGeneric := GenericImplementer{
		name: "Hello Generics",
	}
	processGeneric(convertGenericToString(firstGeneric))
	processGeneric(convertGenericToInt(secondGeneric))
	processGeneric(thirdGeneric)
	fmt.Println("---")
}
