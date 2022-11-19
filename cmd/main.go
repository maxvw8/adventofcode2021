package main

import (
	"advent/pkg/advent"
	"advent/pkg/day1"
	"advent/pkg/day2"
	"fmt"
	"os"
)

const inputPath = "../inputs/"

var problems = []advent.Problem{
	day1.New("1.txt"),
	day2.New(),
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
		sol, err := p.Solve()
		if err != nil {
			fmt.Printf("Error solving problem %s", err)
			os.Exit(1)
		}
		fmt.Printf("[%s] %s\n", p.Name(), sol)
	}

}
