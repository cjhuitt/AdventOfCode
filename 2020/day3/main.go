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

	trees := [5]int{}
	pos := [5]int{}
	move_r := [5]int{1, 3, 5, 7, 1}
	move_d := [5]int{1, 1, 1, 1, 2}
	lineno := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(lineno%move_d[4], lineno%move_d[3])
		for i := 0; i < len(pos); i++ {
			pos[i] = pos[i] % len(line)
			if lineno%move_d[i] == 0 {
				if (line[pos[i]]) == '#' {
					trees[i]++
				}

				pos[i] += move_r[i]
			}
		}
		lineno++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	mult := 1
	for i := 0; i < len(pos); i++ {
		fmt.Println("For slope r", move_r[i], "d", move_d[i], "encountered", trees[i], "trees")
		mult *= trees[i]
	}
	fmt.Println(mult, "trees")
}
