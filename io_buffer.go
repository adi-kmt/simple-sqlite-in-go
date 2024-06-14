package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type inputBuffer struct {
	buffer       *string
	bufferLength int
	inputLength  int
}

func NewInputBuffer() *inputBuffer {
	return &inputBuffer{
		buffer:       nil,
		bufferLength: 0,
		inputLength:  0,
	}
}

func readInput(inputbuffer *inputBuffer) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		return
	}
	input = strings.TrimSuffix(strings.TrimSpace(input), "\n")
	inputbuffer.buffer = &input
	inputbuffer.inputLength = len(input)
}

func closeInputBuffer(inputbuffer *inputBuffer) {
	inputbuffer.buffer = nil
}
