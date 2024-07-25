package main

import (
	"database/sql"
	"fmt"
)

func deleteAllData(db *sql.DB) {

	stmt, errPrepare := db.Prepare("DELETE FROM people WHERE first_name = $1")

	if errPrepare != nil {
		panic(errPrepare)
	}

	defer stmt.Close()

	var errExec error

	_, errExec = stmt.Exec("Java")
	_, errExec = stmt.Exec("Go")

	if errExec != nil {
		panic(errExec)
	}
}

func Lesson4(db *sql.DB) {

	fmt.Printf("Lesson 4 - Delete All Data in DB.")
	deleteAllData(db)
}
