package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	unique_answers := make(map[rune]int)

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			count += len(unique_answers)
			unique_answers = make(map[rune]int)
		} else {
			for _, c := range scanner.Text() {
				unique_answers[c] += 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", count+len(unique_answers), "unique group yes answers")
}
