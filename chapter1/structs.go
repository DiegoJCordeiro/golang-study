package main

import (
	"fmt"
)

type Address struct {
	Street  string
	City    string
	State   string
	Country string
}

type Client struct {
	Name        string
	Age         int32
	Nationality string
	Address     Address
	Admin       bool
}

type Person interface {
	ChangeAdmin()
}

func (client *Client) ChangeAdmin() {
	client.Admin = !client.Admin
	fmt.Printf("%s, Client has changed admin status to %v\n", client.Name, client.Admin)
	fmt.Printf("Client Prototype %+v\n", client)
}

func fillStructs() *Client {

	var firstAddress Address

	firstAddress.Street = "Street"
	firstAddress.City = "City"
	firstAddress.State = "State"
	firstAddress.Country = "Country"

	var firstClient Client
	firstClient.Name = "Person"
	firstClient.Age = 42
	firstClient.Address = firstAddress
	firstClient.Nationality = "Nationality"
	firstClient.Admin = false

	return &firstClient
}

func lesson7() {

	fmt.Println("Lesson 7 - Structs")
	var clientPerson Person = fillStructs()
	clientPerson.ChangeAdmin()
	fmt.Println("---")
}
