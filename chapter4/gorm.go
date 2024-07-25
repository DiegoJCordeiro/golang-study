package main

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func deleteAllDataWithWhere(db *gorm.DB) {

	fmt.Printf("\n-- Gorm -> Delete Many Rows --")

	var people []Person

	db.Where("age = ?", 2).Find(&people)
	db.Delete(people)
}

func updateWithWhere(db *gorm.DB) {

	fmt.Printf("\n-- Gorm -> Update Many Rows --")

	var peoples []Person

	db.Where("people.age = ?", 1).Find(&peoples)

	for _, person := range peoples {
		person.Age = 2
		db.Save(&person)
	}

	db.Where("people.age = ?", 2).Find(&peoples)

	for _, person := range peoples {
		fmt.Printf("%+v \n", person)
	}
}

func queryWithWhere(db *gorm.DB) {

	fmt.Printf("\n-- Gorm -> With Where Clause --\n")

	var people []Person
	db.Where("age=?", "1").Find(&people)

	for _, person := range people {
		fmt.Printf("%+v \n", person)
	}
}

func queryFirstData(db *gorm.DB) {

	fmt.Printf("\n-- Gorm -> First Row --\n")

	var person Person
	db.First(&person, "first_name = ?", "XPTO1")

	fmt.Printf("%+v \n", person)
}

func insertWithGorm(db *gorm.DB) {

	fmt.Printf("\n-- Gorm -> Insert Many Rows --")

	people := []Person{
		{ID: generateUuid(), FirstName: "XPTO1", LastName: "Doe", Age: 1, Description: ""},
		{ID: generateUuid(), FirstName: "XPTO2", LastName: "Doe", Age: 2, Description: ""},
		{ID: generateUuid(), FirstName: "XPTO3", LastName: "Doe", Age: 3, Description: ""},
		{ID: generateUuid(), FirstName: "XPTO4", LastName: "Doe", Age: 4, Description: ""},
		{ID: generateUuid(), FirstName: "XPTO5", LastName: "Doe", Age: 5, Description: ""},
	}

	db.Create(&people)
}

func generateUuid() string {

	id, errUuid := uuid.NewUUID()

	if errUuid != nil {
		panic(errUuid)
	}

	return id.String()
}

func Lesson5(dbConnection *gorm.DB) {

	fmt.Printf("Lesson 5 - Using Gorm.\n")

	insertWithGorm(dbConnection)
	queryFirstData(dbConnection)
	queryWithWhere(dbConnection)
	updateWithWhere(dbConnection)
	deleteAllDataWithWhere(dbConnection)
}
