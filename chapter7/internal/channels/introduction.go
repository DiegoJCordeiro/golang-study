package channels

import "fmt"

var (
	channelMessage chan string
)

func Lesson4() {

	channelMessage = make(chan string)

	go func() {
		channelMessage <- "1 -> Lesson 4 - Intro Channel - Message"
	}()

	message := <-channelMessage
	fmt.Println(message)
}
