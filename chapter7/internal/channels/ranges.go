package channels

import "fmt"

func Lesson6() {

	channelMes := make(chan string)

	go func() {
		for index := 0; index < 10; index++ {
			channelMes <- fmt.Sprintf("%d -> Lesson 6 - Range Channel - Send Message", index)
		}
		close(channelMes)
	}()

	readChannelMessage(channelMes)
}

func readChannelMessage(channelMes chan string) {
	for message := range channelMes {
		fmt.Println(message)
	}
}
