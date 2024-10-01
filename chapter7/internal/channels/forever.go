package channels

import (
	"fmt"
)

func Lesson5() {
	channelMes := make(chan string)

	go func() {
		for index := 0; index < 10; index++ {
			fmt.Printf("%d -> Lesson 5 - Forever Channel - Message \n", index)
		}
		channelMes <- "Done"
	}()
	<-channelMes
}
