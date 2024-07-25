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
	// Address is a composition
	Address Address
	// GeneralRegistry is a composition
	GeneralRegistry GeneralRegistry
	// NaturalPersonsRegister is a composition
	NaturalPersonsRegister NaturalPersonsRegister
	Admin                  bool
}

type Identification struct {
	ID string
}

type GeneralRegistry struct {
	// Identification is an inheritance
	Identification
}

type NaturalPersonsRegister struct {
	// Identification is an inheritance
	Identification
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
	firstClient.GeneralRegistry = GeneralRegistry{
		Identification{
			ID: "1",
		},
	}
	firstClient.NaturalPersonsRegister = NaturalPersonsRegister{
		Identification{
			ID: "2",
		},
	}

	return &firstClient
}

func Lesson7() {

	fmt.Println("Lesson 7 - Structs")
	var clientPerson Person = fillStructs()
	clientPerson.ChangeAdmin()
	fmt.Println("---")
}
