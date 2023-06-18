package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	count := 0
	calories := []int{}
	for _, n := range lines {
		if n == "" {
			calories = append(calories, count)
			count = 0
			continue
		}

		v, err := strconv.Atoi(n)
		if err != nil {
			log.Print(err)
			continue
		}
		count = count + v
	}

	max := 0
	for i, calorie := range calories {
		if i == 0 || calorie > max {
			max = calorie
		}
	}

	// Answer Part 1:
	fmt.Printf("max: %d", max)

	max = 0
	sort.Slice(calories, func(a, b int) bool {
		return calories[b] < calories[a]
	})
	for _, n := range calories[0:3] {
		max = max + n
	}

	// Answer Part 2:
	fmt.Printf("max: %d", max)
}
