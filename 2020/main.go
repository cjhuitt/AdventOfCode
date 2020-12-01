package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	entries := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		for _, s := range values {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			entries = append(entries, i)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(entries)

	i := 0
	j := len(entries) - 1
	for true {
		val := entries[i] + entries[j]
		if val == 2020 {
			break
		} else if val > 2020 {
			j--
		} else if val < 2020 {
			i++
		}
	}

	fmt.Println(entries[i] * entries[j])
}
