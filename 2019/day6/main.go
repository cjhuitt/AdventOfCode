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

	specs := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		specs = append(specs, scanner.Text())
	}

	fmt.Println("Found", len(specs), "specifications")
}
