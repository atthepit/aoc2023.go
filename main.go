package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/atthepit/aoc2023/solutions"
)

func main() {
	day := flag.Int("day", 1, "Day to execute")
	part := flag.String("part", "a", "Day part to execute")
	filepath := flag.String("file", "", "Path to the file with the input for the task")
	flag.Parse()

	reader, err := os.OpenFile(*filepath, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err.Error())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(reader)

	if *day == 1 {
		solution, err := solutions.Day01(scanner, *part)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(solution)
		os.Exit(0)
	}

	fmt.Printf("Unkown params\n")
	os.Exit(1)
}
