package main

import (
	"errors"
	"fmt"
)

// Anonymous Function - Sum Two Values - Return total and error
func sumWithAnonymousFunctions(firstValue int, secondValue int) (total int, err error) {

	total, err = func() (int, error) {
		return sumOnlyTwoValues(firstValue, secondValue)
	}()

	return
}

// Default Function - Sum Two Values - Return total and error
func sumOnlyTwoValues(firstValue int, secondValue int) (int, error) {

	valueSummed := firstValue + secondValue

	if valueSummed > 10 {
		return valueSummed, errors.New("error value is major that 10")
	}

	return valueSummed, nil
}

// Variadic Function - Sum Many Values - Return total and error
func sumManyValues(values ...int) (int, error) {

	var total = 0
	for _, value := range values {
		total += value
	}

	return total, nil
}

// Lesson6() is about functions variadic, default and returns.
func Lesson6() {

	fmt.Println("Lesson 6 - Functions")

	valueSummed, err1 := sumOnlyTwoValues(1, 10)
	valueSummedManyValues, err2 := sumManyValues(1, 2, 3, 4, 5)
	valueSummedWithAnonymousFunctions, err3 := sumWithAnonymousFunctions(1, 9)

	fmt.Printf("Total = %d, Error = %v\n", valueSummed, err1)
	fmt.Printf("Total = %d, Error = %v\n", valueSummedManyValues, err2)
	fmt.Printf("Total = %d, Error = %v\n", valueSummedWithAnonymousFunctions, err3)

	fmt.Println("---")
}
