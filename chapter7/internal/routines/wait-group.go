package routines

import (
	"fmt"
	"sync"
	"time"
)

func Lesson2() {

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go taskLesson2("Lesson 2 - WaitGroup - Thread 1", &waitGroup)
	go taskLesson2("Lesson 2 - WaitGroup - Thread 2", &waitGroup)

	waitGroup.Wait()
}

func taskLesson2(name string, waitGroup *sync.WaitGroup) {

	for index := 0; index < 10; index++ {
		fmt.Printf("%d -> %s is running \n", index, name)
		time.Sleep(5 * time.Second)
	}

	waitGroup.Done()
}
