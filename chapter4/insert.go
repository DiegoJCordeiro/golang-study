package main

import (
	"database/sql"
	"fmt"
)

func insertData(db *sql.DB, person *Person) {
	stmt, errPrepare := db.Prepare("INSERT INTO people(id, first_name, last_name, age, description) VALUES($1, $2, $3, $4, $5)")
	if errPrepare != nil {
		panic(errPrepare)
	}
	defer stmt.Close()
	_, errExec := stmt.Exec(person.ID, person.FirstName, person.LastName, person.Age, person.Description)
	if errExec != nil {
		panic(errExec)
	}
}

func Lesson1(db *sql.DB) {

	fmt.Printf("Lesson 1 - Insert Data in DB.")

	var firstPerson Person
	var secondPerson Person

	firstPerson.NewPerson("Go", "", 1, "Golang.")
	secondPerson.NewPerson("Java", "", 2, "Java.")

	insertData(db, &firstPerson)
	insertData(db, &secondPerson)
}
