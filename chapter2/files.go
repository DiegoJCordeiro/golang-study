package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	errWrite  error
	errClose  error
	errFlush  error
	errOpen   error
	errCreate error
)

func createFile() {

	file, errCreate := os.Create("./files/test.txt")

	if errCreate != nil {
		panic(errCreate)
	}

	errClose := file.Close()

	if errClose != nil {
		panic(errClose)
	}
}

func writeFile() {

	errWrite := os.WriteFile("./test.txt", []byte("- First Line here."), 777)

	if errWrite != nil {
		panic(errWrite)
	}
}

func writeLineToLineFile() {
	file, errOpen := os.OpenFile("./files/test.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if errOpen != nil {
		panic(errOpen)
	}

	writer := bufio.NewWriter(file)

	_, errWrite = writer.WriteString("\n")
	_, errWrite = writer.Write([]byte("- Second Line."))
	_, errWrite = writer.WriteString("\n")
	_, errWrite = writer.Write([]byte("- Third Line."))
	_, errWrite = writer.WriteString("\n")
	_, errWrite = writer.Write([]byte("- Fourth Line."))
	errFlush := writer.Flush()

	errClose := file.Close()

	if errClose != nil {
		panic(errClose)
	} else if errFlush != nil {
		panic(errFlush)
	} else if errWrite != nil {
		panic(errWrite)
	}
}

func readFile() {

	file, errOpen := os.ReadFile("./files/test.txt")

	if errOpen != nil {
		panic(errOpen)
	}

	fmt.Printf("Content file: \n%s \n", string(file))
}

func readLineToLineFile() {

	file, errOpen := os.Open("./files/test.txt")

	if errOpen != nil {
		panic(errOpen)
	}

	reader := bufio.NewReader(file)

	fmt.Printf("Content file: \n")
	for {
		line, _, errReader := reader.ReadLine()

		if errReader != nil {
			fmt.Printf("End Content.\n")
			break
		}

		fmt.Printf("%s\n", string(line))
	}
}

func Lesson1() {

	fmt.Println("Lesson 1 - Manipulate Files")

	createFile()
	writeFile()
	readFile()
	writeLineToLineFile()
	readLineToLineFile()

	fmt.Println("---")
}
