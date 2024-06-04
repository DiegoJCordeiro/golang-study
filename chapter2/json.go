package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Account struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}

func CreateAccount(name string, createdBy string, createdAt time.Time) Account {

	var newAccount = Account{
		Id:        1,
		Name:      name,
		CreatedBy: createdBy,
		CreatedAt: createdAt,
	}

	return newAccount
}

func marshalAccount(account Account) []byte {

	res, errorMarshal := json.Marshal(account)

	if errorMarshal != nil {
		panic(errorMarshal)
	}

	return res
}

func unmarshalAccount(account []byte) Account {

	var newAccount Account

	errorUnmarshal := json.Unmarshal(account, &newAccount)

	if errorUnmarshal != nil {
		panic(errorUnmarshal)
	}

	return newAccount
}

func encodeAccount(account Account) {

	errorEncoder := json.NewEncoder(os.Stdout).Encode(account)

	if errorEncoder != nil {
		panic(errorEncoder)
	}
}

func Lesson4() {

	fmt.Println("Lesson 4 - Manipulate Json")

	newAccountToMarshal := CreateAccount("GoLang", "Diego", time.Time{})
	fmt.Printf("Marshal Json: %s \n", string(marshalAccount(newAccountToMarshal)))
	fmt.Printf("Encoder Json: ")
	encodeAccount(newAccountToMarshal)
	newAccountToUnmarshal := unmarshalAccount([]byte(`{"id":1, "name":"GoLang","createdBy":"Diego","createdAt":"0001-01-01T00:00:00Z"}`))
	fmt.Printf("Unmarshal Json: %+v \n", newAccountToUnmarshal)
}
