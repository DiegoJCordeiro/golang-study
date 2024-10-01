package routines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	counterAtomic int64 = 0
	counter       int64 = 0
	mutex               = sync.Mutex{}
)

func Lesson3() {

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go taskLesson3("Lesson 3 - Mutex - Thread 1", &waitGroup)
	go taskLesson3("Lesson 3 - Mutex - Thread 2", &waitGroup)

	waitGroup.Wait()

	fmt.Printf("* -> Lesson 3 - Mutex  - Value = %d \n", counter)
	fmt.Printf("* -> Lesson 3 - Atomic - Value = %d \n", counter)
}

func taskLesson3(name string, waitGroup *sync.WaitGroup) {

	for index := 0; index < 10; index++ {
		fmt.Printf("%d -> %s is running \n", index, name)
		mutex.Lock()
		counter++
		mutex.Unlock()
		atomic.AddInt64(&counterAtomic, 1)
		time.Sleep(5 * time.Second)
	}

	waitGroup.Done()
}
