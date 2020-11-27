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

	// Attempt to find proper execution input
	value := 0
	i := 0
	j := -1
	for value != 19690720 {
		temp := make([]int, len(stack))
		copy(temp, stack)

		temp[1] = i
		temp[2] = j

		program := Intcode.New(temp)
		program = program.Execute()
		value = program.Data()[0]
		if value == 19690720 {
			break
		}
		j++
		if j >= 100 {
			i++
			if i >= 100 {
				break
			}
			j = 0
		}
	}

	fmt.Println("100 *", i, "+", j, "=", (100*i)+j)
}
