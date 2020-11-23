package main

import (
    "bufio"
    "fmt"
    "log"
    "os"

//    "SpacecraftModule"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err:= scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
