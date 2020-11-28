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

	program := Intcode.New(stack)
	input := 5
	for !program.IsDone() {
		program = program.Execute()
		if program.WantsInput() {
			fmt.Println("In:", input)
			program = program.WithInput(&input)
		}
		if program.HasOutput() {
			fmt.Println("Out:", *program.Output())
		}
	}

	fmt.Println("Done")
}
