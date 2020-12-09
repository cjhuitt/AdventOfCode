package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"bags"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	specs := bags.ParseSpecs(lines)
	total := 0
	for _, spec := range specs {
		if spec.TotalAllowed("shiny gold", specs) > 0 {
			total++
		}
	}

	fmt.Println("Total allowing shiny gold:", total)
}
