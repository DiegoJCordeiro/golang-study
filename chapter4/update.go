package main

import (
	"database/sql"
	"fmt"
)

func updateGoLang(db *sql.DB) {

	stmt, errPrepare := db.Prepare("UPDATE people SET description = $1 WHERE first_name = $2")

	if errPrepare != nil {
		panic(errPrepare)
	}

	defer stmt.Close()

	_, errExec := stmt.Exec("Go - Program Language", "Go")

	if errExec != nil {
		panic(errExec)
	}
}

func updateJavaLang(db *sql.DB) {

	stmt, errPrepare := db.Prepare("UPDATE people SET description = $1 WHERE first_name = $2")

	if errPrepare != nil {
		panic(errPrepare)
	}

	defer stmt.Close()

	_, errExec := stmt.Exec("Java - Program Language", "Java")

	if errExec != nil {
		panic(errExec)
	}
}

func Lesson2(db *sql.DB) {

	fmt.Printf("Lesson 2 - Update Data in DB.")

	updateGoLang(db)
	updateJavaLang(db)
}
