package main

import (
	"fmt"
	"strings"
)

func main() {
	input := []string{"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw"}

    priorityMap := map[string]int{"p": 16, "s": 19, "t": 20, "v": 22, "L": 38, "P": 42}

    count := 0
	for _, rucksack := range input {
		rucksackItems := strings.Split(rucksack, "")
        mid := len(rucksackItems) / 2;
        compartment1 := rucksackItems[:mid]
        compartment2 := rucksackItems[mid:]
        fmt.Println("Mid: ", mid)
        fmt.Println("Compartment 1: ",compartment1)
        fmt.Println("Compartment 2: ", compartment2)

        match := ""
        for _, c1Item := range compartment1 {
            for _, c2Item := range compartment2 {
                if c1Item == c2Item {
                    match = c1Item
                }
            }
        }

        fmt.Printf("Match: %v(%d)\n\n", match, priorityMap[match])
        count = count + priorityMap[match]
	}

    fmt.Println("Total: ", count)

}
