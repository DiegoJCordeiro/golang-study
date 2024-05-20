package main

import (
	"fmt"
)

type PhysicalPerson struct {
	fit string
}

type TypeFit interface {
	changeFit(fitType string)
	getFit() string
}

func (physicalPerson *PhysicalPerson) changeFit(fitType string) {
	physicalPerson.fit = fitType
}
func (physicalPerson *PhysicalPerson) getFit() string {
	return physicalPerson.fit
}

func showFit(typeFit TypeFit) {
	fmt.Printf("Fit Type is: %v\n", typeFit.getFit())
}

func lesson9() {

	fmt.Println("Lesson 9 - Interfaces")
	var physicalPerson = PhysicalPerson{
		fit: "Ecto morph",
	}
	physicalPerson.changeFit("Mesomorph")
	showFit(&physicalPerson)
	fmt.Println("---")
}
