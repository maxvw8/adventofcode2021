package main

import (
	"advent/pkg/part1"
	"advent/pkg/utils"
	"fmt"
	"os"
)

const inputPath = "../inputs/"

var problems = []utils.Problem{
	part1.New("1.txt"),
}

func main() {
	for _, p := range problems {
		file, err := os.Open(fmt.Sprintf("%s/%s", inputPath, p.FileName()))
		if err != nil {
			fmt.Printf("Error while open file %s", err)
			os.Exit(1)
		}
		err = p.Load(file)
		if err != nil {
			fmt.Printf("Error parsing file %s", err)
			os.Exit(1)
		}
		fmt.Printf("Result is %s", p.Solve())
	}

}
