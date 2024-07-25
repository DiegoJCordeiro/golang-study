package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	waitGroup                sync.WaitGroup
	syncMutex                sync.Mutex
	numberCounterOneThousand int
)

func showMessage(wait *sync.WaitGroup) {

	fmt.Printf(" \n--- START - showMessage function 10 ---\n")
	for counter := 10; counter >= 0; counter-- {
		counterConv := strconv.Itoa(counter)
		fmt.Printf("- %s \n", counterConv)
	}
	fmt.Printf("--- END - showMessage function 10 ---\n")
	wait.Done()
}

func skipIndexUntilTen(wait *sync.WaitGroup) {
	fmt.Printf(" \n--- START - skip function until 10 ---\n")
	syncMutex.Lock()
	for counter := 0; counter < 10; counter++ {
		numberCounterOneThousand++
		fmt.Printf("- it is in %d until 10 \n", numberCounterOneThousand)
	}
	syncMutex.Unlock()
	fmt.Printf(" --- END - skip function until 10 ---\n")
	wait.Done()
}

func skipIndexUntilFifth(wait *sync.WaitGroup) {
	fmt.Printf(" \n--- START - skip function until 50 ---\n")
	syncMutex.Lock()
	for counter := 0; counter < 50; counter++ {
		numberCounterOneThousand++
		fmt.Printf("- it is in %d until 50 \n", numberCounterOneThousand)
	}
	syncMutex.Unlock()
	fmt.Printf(" --- END - skip function until 50 ---\n")
	wait.Done()
}

func Lesson2() {

	fmt.Printf("Lesson 2 - Concurrency\n")

	waitGroup.Add(3)

	go showMessage(&waitGroup)
	go skipIndexUntilTen(&waitGroup)
	go skipIndexUntilFifth(&waitGroup)

	time.Sleep(5 * time.Second)
}
