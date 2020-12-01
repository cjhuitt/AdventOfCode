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

	val := -1
	i := -1
	j := -1
	k := -1
done:
	for i = 0; i < len(entries) && val != 2020; i++ {
		for j = i + 1; j < len(entries) && val != 2020; j++ {
			for k = j + 1; k < len(entries) && val != 2020; k++ {
				val = entries[i] + entries[j] + entries[k]
				if val == 2020 {
					break done
				}
			}
		}
	}

	fmt.Println(i, j, k, entries[i], entries[j], entries[k])
	fmt.Println(entries[i] * entries[j] * entries[k])
}
