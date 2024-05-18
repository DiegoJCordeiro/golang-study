package main

import (
	"fmt"
)

// Declaration Global Constant Variable
const CONSTANT_VARIABLE_STRING string = "CONSTANT_VARIABLE_STRING"

// User - Creation Variable Type
type User struct {
	firstName string
	lastName  string
	age       int32
}

// Declaration Too Much Global Variables
var (
	// alias for int32 (0 to 4294967295)
	integerVariable int

	// 8-bit integers (0 to 255)
	uint8Variable uint8

	// 16-bit integers (0 to 65535)
	uint16Variable uint16

	// 32-bit integers (0 to 4294967295)
	uint32Variable uint32

	// 64-bit integers (0 to 18446744073709551615)
	uint64Variable uint64

	// 8-bit integers (-128 to 127)
	int8Variable int8

	// 16-bit integers (-32768 to 32767)
	int16Variable int16

	// 32-bit integers (-2147483648 to 2147483647)
	int32Variable int32

	// 32-bit floating-point numbers
	float32Variable float32

	// 64-bit floating-point numbers
	float64Variable float64

	// Complex numbers with float32 real and imaginary parts
	complex64Variable complex64

	// Complex numbers with float64 real and imaginary parts
	complex128Variable complex128

	// Boolean True or False
	booleanVariable bool

	// strings of 8-bit bytes. Values of string type are immutable.
	stringVariable string

	// alias for uint8 rune alias for int32
	byteVariable byte

	user User
)

func init() {
	fillVariables()
}

func fillVariables() {

	integerVariable = 4294967295
	uint8Variable = 255
	uint16Variable = 65535
	uint32Variable = 4294967295
	uint64Variable = 18446744073709551615
	int8Variable = 127
	int16Variable = 32767
	int32Variable = 2147483647

	float32Variable = 1.0
	float64Variable = 1.0
	complex64Variable = 1.0
	complex128Variable = 1.0

	booleanVariable = true
	stringVariable = "stringVariable"

	user = User{
		firstName: "Study",
		lastName:  "Go",
		age:       1,
	}

	byteVariable = uint8Variable
}

func showVariables() {

	// := Apply and set the value and variable type.
	localVariable := "localVariable"
	fmt.Printf("(Type %T) Local Variable - Value: %s\n", localVariable, localVariable)
	fmt.Printf("(Type %T) Constant Variable - Value: %s\n", CONSTANT_VARIABLE_STRING, CONSTANT_VARIABLE_STRING)
	fmt.Printf("(Type %T) Variable - Value: %d\n", integerVariable, integerVariable)
	fmt.Printf("(Type %T) Variable - Value: %d\n", uint8Variable, uint8Variable)
	fmt.Printf("(Type %T) Variable - Value: %d\n", uint16Variable, uint16Variable)
	fmt.Printf("(Type %T) Variable - Value: %d\n", uint32Variable, uint32Variable)
	fmt.Printf("(Type %T) Variable - Value: %d\n", uint64Variable, uint64Variable)
	fmt.Printf("(Type %T) Variable - Value: %d\n", int8Variable, int8Variable)
	fmt.Printf("(Type %T) Variable - Value: %d\n", int16Variable, int16Variable)
	fmt.Printf("(Type %T) Variable - Value: %d\n", int32Variable, int32Variable)
	fmt.Printf("(Type %T) Variable - Value: %f\n", float32Variable, float32Variable)
	fmt.Printf("(Type %T) Variable - Value: %f\n", float64Variable, float64Variable)
	fmt.Printf("(Type %T) Variable - Value: %f\n", complex64Variable, complex64Variable)
	fmt.Printf("(Type %T) Variable - Value: %f\n", complex128Variable, complex128Variable)
	fmt.Printf("(Type %T) Variable - Value: %t\n", booleanVariable, booleanVariable)
	fmt.Printf("(Type %T) Variable - Value: %s\n", stringVariable, stringVariable)
	fmt.Printf("(Type %T) Variable - Value: %v", user, user)
	fmt.Printf("(Type %T) Variable - Value: %d\n", byteVariable, byteVariable)
}

// lesson2() is about imports and variables.
func lesson2() {

	fmt.Println("Lesson 2 - Variables")
	showVariables()
	fmt.Println("---")
}
