package main

import (
	"advent/pkg/part1"
	"advent/pkg/utils"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../inputs/1.txt")
	if err != nil {
		fmt.Printf("Error while open file %s", err)
		os.Exit(1)
	}
	input, err := utils.ReadFile(file, strconv.Atoi)
	if err != nil {
		fmt.Printf("Error parsing file %s", err)
		os.Exit(1)
	}
	output := part1.TimesIncreased(input)
	fmt.Printf("Result is: %d", output)
}
