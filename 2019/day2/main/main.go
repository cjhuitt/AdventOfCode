package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"Intcode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stack := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		for _, s := range values {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			stack = append(stack, i)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Adapt to last good state
	stack[1] = 12
	stack[2] = 2

	// Go
	program := Intcode.New(stack)
	program = program.Execute()

	fmt.Println("First position value:", program.Data()[0])
}
