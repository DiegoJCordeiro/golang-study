package routines

import (
	"fmt"
	"time"
)

func Lesson1() {

	go taskLesson1("Lesson 1 - Intro Goroutine - Thread 1")
	go taskLesson1("Lesson 1 - Intro Goroutine - Thread 2")
	go func(name string) {
		for index := 0; index < 10; index++ {
			fmt.Printf("%d -> %s is running \n", index, name)
			time.Sleep(5 * time.Second)
		}
	}("Lesson 1 - Intro Goroutine - Thread 3")

	time.Sleep(40 * time.Second)
}

func taskLesson1(name string) {

	for index := 0; index < 10; index++ {
		fmt.Printf("%d -> %s is running \n", index, name)
		time.Sleep(5 * time.Second)
	}
}
