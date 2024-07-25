package main

import (
	"database/sql"
	"fmt"
)

func queryData(db *sql.DB) {

	fmt.Printf("--Query One Row--")

	stmt, errPrepare := db.Prepare("SELECT id, first_name, last_name, age, description FROM people WHERE first_name LIKE $1")

	if errPrepare != nil {
		panic(errPrepare)
	}

	var person Person

	errScan := stmt.QueryRow("Go").Scan(&person.ID, &person.FirstName, &person.LastName, &person.Age, &person.Description)

	if errScan != nil {
		panic(errScan)
	}

	fmt.Printf("\n%+v", person)
}

func queryAllData(db *sql.DB) {

	fmt.Printf("\n--Query Many Rows--")

	stmt, errQuery := db.Query("SELECT id, first_name, last_name, age, description FROM people")

	if errQuery != nil {
		panic(errQuery)
	}

	defer stmt.Close()

	for stmt.Next() {
		var personFound Person

		errScan := stmt.Scan(&personFound.ID, &personFound.FirstName, &personFound.LastName, &personFound.Age, &personFound.Description)

		if errScan != nil {
			panic(errScan)
		}

		fmt.Printf("\n%+v", personFound)
	}
}

func Lesson3(db *sql.DB) {

	fmt.Printf("Lesson 3 - Query Data in DB.\n")
	queryData(db)
	queryAllData(db)
}
