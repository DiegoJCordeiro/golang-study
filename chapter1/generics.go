package main

import (
	"fmt"
	"reflect"
)

var (
	firstGeneric  string
	secondGeneric int32
)

type GenericImplementer struct {
	name string
}

type GenericConstraint interface {
	~int32 | ~string | GenericImplementer
}

func processGeneric[T GenericConstraint](generic T) {
	fmt.Printf("Processing generic %T: %+v\n", generic, generic)
}

// Function return Type Assertion
func convertGeneric[T GenericConstraint](generic T) T {

	isInteger := reflect.TypeOf(generic) == reflect.TypeOf((*int)(nil))

	if isInteger {
		return generic
	} else {
		return generic
	}

}

func Lesson10() {

	fmt.Println("Lesson 10 - Generics")
	firstGeneric = "Hello Generics"
	secondGeneric = 0
	thirdGeneric := GenericImplementer{
		name: "Hello Generics",
	}
	processGeneric(convertGeneric(firstGeneric))
	processGeneric(convertGeneric(secondGeneric))
	processGeneric(thirdGeneric)
	fmt.Println("---")
}
