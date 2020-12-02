package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"pwparser"
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

	t1 := 0
	t2 := 0
	for _, c := range candidates {
		parts := strings.Split(c, ":")
		if len(parts) != 2 {
			log.Fatal("Bad pw entry")
		}

		r, err := pwparser.ParseRule(parts[0])
		if err != nil {
			log.Fatal(err)
		}

		pw := strings.TrimSpace(parts[1])
		if r.MatchesType1(pw) {
			t1++
		}
		if r.MatchesType2(pw) {
			t2++
		}
	}

	fmt.Println("Found", t1, "type 1 matches")
	fmt.Println("Found", t2, "type 2 matches")
}
