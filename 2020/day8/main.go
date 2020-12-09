package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"vm"
)

func main() {
	file, err := os.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	source := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		source = append(source, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	orig := vm.Parse(source)
	_, program := orig.Execute()
	fmt.Println("Accumulator at first repeat instruction is", program.Accumulator())

	program = orig
	test, program, found := orig.FixNextNop(0)
	fmt.Println("Next", test)
	good := false
	last := orig
	for found && !good {
		good, last = program.Execute()
		test, program, found = orig.FixNextNop(test + 1)
		fmt.Println("Next", test)
	}

	test, program, found = orig.FixNextJmp(0)
	fmt.Println("Next", test)
	for found && !good {
		good, last = program.Execute()
		test, program, found = orig.FixNextJmp(test + 1)
		fmt.Println("Next", test)
	}

	if good {
		fmt.Println("After fix, accumulator is", last.Accumulator())
	}
}
