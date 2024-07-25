package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Person struct {
	ID          string `gorm:"primaryKey"`
	FirstName   string
	LastName    string
	Age         int
	Description string
	gorm.Model
}

type CreatePerson interface {
	NewPerson(firstName string, lastName string, age int, description string)
}

func (person *Person) NewPerson(firstName string, lastName string, age int, description string) {

	id, _ := uuid.NewUUID()
	person.ID = id.String()
	person.FirstName = firstName
	person.LastName = lastName
	person.Age = age
	person.Description = description
}

func main() {

	dbSQL := openConnectionWithDatabase("postgres")
	dbGORM := openConnectionWithDatabaseGORM("postgres")

	Lesson1(dbSQL)
	breakLine()
	Lesson2(dbSQL)
	breakLine()
	Lesson3(dbSQL)
	breakLine()
	Lesson4(dbSQL)
	breakLine()
	Lesson5(dbGORM)
	breakLine()

	defer dbSQL.Close()
}

func openConnectionWithDatabase(dbName string) *sql.DB {

	db, err := sql.Open("postgres", "postgres://postgres:1234@localhost:5432/"+dbName+"?sslmode=disable")

	if err != nil {
		panic(err)
	}

	return db
}

func openConnectionWithDatabaseGORM(dbName string) *gorm.DB {

	datasourceName := "postgres://postgres:1234@localhost:5432/" + dbName + "?sslmode=disable"

	db, errOpen := gorm.Open(postgres.Open(datasourceName), &gorm.Config{})

	if errOpen != nil {
		panic(errOpen)
	}

	errAutoMigrate := db.AutoMigrate(&Person{})

	if errAutoMigrate != nil {
		panic(errOpen)
	}

	return db
}

func breakLine() {

	fmt.Printf("\n---\n")
}
