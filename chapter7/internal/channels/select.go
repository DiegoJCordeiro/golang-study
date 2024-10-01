package channels

import (
	"fmt"
	"time"
)

func Lesson8() {

	channelOne := make(chan int, 2)
	channelTwo := make(chan int, 2)

	go func() {
		channelOne <- 1
		time.Sleep(2 * time.Second)
	}()

	go func() {
		channelTwo <- 3
		time.Sleep(3 * time.Second)
	}()

	select {
	case <-channelOne:
		channelSelected := fmt.Sprintf("%d -> Lesson 8 - Select Channel - Listen Message", 1)
		fmt.Println(channelSelected)
	case <-channelTwo:
		channelSelected := fmt.Sprintf("%d -> Lesson 8 - Select Channel - Listen Message", 2)
		fmt.Println(channelSelected)
	case <-time.After(time.Second * 3):
		channelTimeout := fmt.Sprintf("%d -> Lesson 8 - Select Channel - Listen Message", 3)
		fmt.Println(channelTimeout)
	}
}
