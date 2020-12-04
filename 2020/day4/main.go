package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"passports"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	valid := 0
	item := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			item += " " + line
		} else {
			if passports.Parse(item).IsValid() {
				valid++
			}
			item = ""
			total++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if item != "" {
		if passports.Parse(item).IsValid() {
			valid++
		}
		item = ""
		total++
	}

	fmt.Println(valid, "of", total, "are valid")
}
