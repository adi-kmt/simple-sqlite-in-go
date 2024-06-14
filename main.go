package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		inputBuffer := NewInputBuffer()
		for {
			printPrompt()
			readInput(inputBuffer)

			if *inputBuffer.buffer == ".exit" {
				fmt.Println("Bye!")
				closeInputBuffer(inputBuffer)
				os.Exit(0)
			} else {
				fmt.Printf("Unrecognized command '%s'.\n", *inputBuffer.buffer)
			}
		}
	}
}
