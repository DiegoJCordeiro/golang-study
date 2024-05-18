package main

import (
	"fmt"
)

var (
	mapLocationVariable = map[string]string{}
	mapSkillsVariable   = make(map[string]string)
)

func init() {
	fillMapVariable()
}

func fillMapVariable() {
	mapLocationVariable["Country"] = "USA"
	mapLocationVariable["State"] = "California"
	mapLocationVariable["City"] = "Los Angeles"

	mapSkillsVariable["Backend"] = "Java"
	mapSkillsVariable["Frontend"] = "Angular | Reactjs"
	mapSkillsVariable["Database"] = "Postgres | Oracle"
}

func manipulateMapVariable() {

	delete(mapLocationVariable, "Country")
	mapLocationVariable["Country"] = "United States"

	delete(mapSkillsVariable, "Backend")
	mapSkillsVariable["Backend"] = "Go"
}

func showMapVariable() {

	fmt.Println("Location")
	for keyLocation, valueLocation := range mapLocationVariable {
		fmt.Printf("Key = %s, Value = %s\n", keyLocation, valueLocation)
	}

	fmt.Println("Skills")
	for keySkill, valueSkill := range mapSkillsVariable {
		fmt.Printf("Key = %s, Value = %s\n", keySkill, valueSkill)
	}
}

func lesson5() {

	fmt.Println("Lesson 5 - Maps")
	manipulateMapVariable()
	showMapVariable()
	fmt.Println("---")
}
