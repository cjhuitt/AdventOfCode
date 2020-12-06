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

	answers := make(map[rune]int)

	any_count := 0
	all_count := 0
	members := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			for _, c := range answers {
				if c == members {
					all_count++
				}
				any_count++
			}
			answers = make(map[rune]int)
			members = 0
		} else {
			for _, c := range scanner.Text() {
				answers[c] += 1
			}
			members++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", any_count, "unique group yes answers")
	fmt.Println("Part 1:", all_count, "complete group yes answers")
}
