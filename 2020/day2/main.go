package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("passwords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	candidates := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		candidates = append(candidates, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Candidates:", len(candidates))
}
