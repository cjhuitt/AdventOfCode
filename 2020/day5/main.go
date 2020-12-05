package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"seats"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	max := -1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		id := seats.Find(scanner.Text()).Id()
		if id > max {
			max = id
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Max id:", max)
}
