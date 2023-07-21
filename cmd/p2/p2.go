package main

import (
	"fmt"
	"os"
	"strings"
)

func main () {
    dat, err := os.ReadFile("input.txt") 
    if err != nil {
        panic(err)
    }

    lines := strings.Split(string(dat), "\n")
    fmt.Println(lines)

    for _, line := range lines {
        fmt.Println(line)
    }
}

