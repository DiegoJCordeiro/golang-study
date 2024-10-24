package main

import (
	"fmt"
	"github.com/DiegoJCordeiro/golang-private-repositories/pkg/events"
)

func main() {
	var eventDispatcher events.EventDispatcher = events.NewEventDispatcher()
	fmt.Println(eventDispatcher)
}
