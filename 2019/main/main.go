package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"

//    "SpacecraftModule"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    masses := make([]int, 0)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        i, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }
        masses = append(masses, i)
    }

    if err:= scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
