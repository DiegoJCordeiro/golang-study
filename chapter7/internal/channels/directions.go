package channels

import (
	"fmt"
	"time"
)

func Lesson7() {

	channel := make(chan string, 10)

	go func() {
		for index := 0; index <= 10; index++ {
			fillChannel(fmt.Sprintf("%d -> Lesson 7 - Direction Channel - Send Message", index), channel)
			go readChannel(channel)
		}
		close(channel)
	}()

	time.Sleep(time.Second * 5)
}

func fillChannel(message string, channel chan<- string) {
	channel <- message
}

func readChannel(channel <-chan string) {
	fmt.Println(<-channel)
}
