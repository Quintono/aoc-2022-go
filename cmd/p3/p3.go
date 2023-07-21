package main

import (
	"fmt"
	"os"
	"strings"
)

var priorityMap = map[string]int{}

func initializePriorityMap() {
	for i := 0; i < 26; i++ {
		priorityMap[string(rune('a'+i))] = i + 1
		priorityMap[string(rune('A'+i))] = i + 27
	}
}

func solvePartOne(lines []string) {
	count := 0
	for _, rucksack := range lines {
		rucksackItems := strings.Split(rucksack, "")
		mid := len(rucksackItems) / 2
		compartment1 := rucksackItems[:mid]
		compartment2 := rucksackItems[mid:]
		fmt.Println("Mid: ", mid)
		fmt.Println("Compartment 1: ", compartment1)
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

func solvePartTwo(lines []string) {
	count := 0
	var groups [][]string

	for i := 3; i <= len(lines); i += 3 {
		groups = append(groups, lines[i-3:i])
	}

	fmt.Println("groups: ", groups)

	for _, group := range groups {
		var elf1 = group[0]
		var elf2 = group[1]
		var elf3 = group[2]
		match := ""

	exit:
		for _, elf1Item := range elf1 {
			for _, elf2Item := range elf2 {
				if elf1Item == elf2Item {
					for _, elf3Item := range elf3 {
						if elf1Item == elf3Item {
							match = string(elf1Item)
							count = count + priorityMap[match]
							break exit
						}
					}
				}
			}
		}
	}

	fmt.Println("Total: ", count)
}

func main() {
	initializePriorityMap()

	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	//fmt.Println("Lines: ", lines)

	//solvePartOne(lines)
	solvePartTwo(lines)
}

