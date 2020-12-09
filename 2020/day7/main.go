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
	may_contain := 0
	for _, spec := range specs {
		if spec.TotalAllowed("shiny gold", specs) > 0 {
			may_contain++
		}
	}

	fmt.Println("Total that may contain shiny gold:", may_contain)

	g := specs["shiny gold"]
	fmt.Println("Total bags necessary in shiny gold:", g.TotalContained(specs))
}
