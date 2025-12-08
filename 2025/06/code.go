package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"strings"
	"strconv"
	"log"
)

func main() {
	aoc.Harness(run)
}

// Create an enum of possible operations
type Operation int

const (
	Add Operation = iota
	Multiply
)

func parseInput(fname string) (numbers [][]int, operations []Operation) {
	lines := strings.Split(strings.TrimSpace(fname), "\n")

	for _, line := range lines[:len(lines)-1] {
		values := strings.Fields(line)
		row := make([]int, len(values))
		for i, value := range values {
			row[i], _ = strconv.Atoi(value)
		}
		numbers = append(numbers, row)
	}

	for _, op := range strings.Fields(lines[len(lines)-1]) {
		switch op {
		case "+":
			operations = append(operations, Add)
		case "*":
			operations = append(operations, Multiply)
		}
	}

	return	
}

func parseInputPart2(fname string) (numbers [][]int, operations []Operation) {
	lines := strings.Split(strings.TrimSpace(fname), "\n")
	line_length := 0
	for _, line := range lines[:len(lines)-1] {
		if len(line) > line_length {
			line_length = len(line)
		}
	}
	numbers = append(numbers, []int{})
	for j := 0; j < line_length; j++ {
		currentColumn := []byte{}
		for _, line := range lines[:len(lines)-1] {
			if j >= len(line) {
				continue
			}
			if line[j] != ' ' {
				currentColumn = append(currentColumn, line[j])
			}
		}

		if len(currentColumn) == 0{
			numbers = append(numbers, []int{})
			continue
		}
		intValue, err := strconv.Atoi(string(currentColumn))
		if err != nil {
			log.Fatalf("invalid number in column %d: %s", j, string(currentColumn))
		}
		numbers[len(numbers)-1] = append(numbers[len(numbers)-1], intValue)
	}
	for _, op := range strings.Fields(lines[len(lines)-1]) {
		switch op {
		case "+":
			operations = append(operations, Add)
		case "*":
			operations = append(operations, Multiply)
		}
	}
	return	
}


// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	var numbers [][]int
	var operations []Operation
	grandTotal := 0
	if part2 {
		numbers, operations = parseInputPart2(input)
		for i := 0; i < len(operations); i++ {
			colTotal := 0
			for value := range numbers[i] {
				switch operations[i] {
				case Add:
					colTotal += numbers[i][value]
				case Multiply:
					if colTotal == 0 {
						colTotal = 1
					}
					colTotal *= numbers[i][value]
				}
			}
			grandTotal += colTotal
		}
	} else {
		numbers, operations = parseInput(input)
		for j := 0; j < len(operations); j++ {
			colTotal := 0
			for i := 0; i < len(numbers); i++ {
				switch operations[j] {
				case Add:
					colTotal += numbers[i][j]
				case Multiply:
					if colTotal == 0 {
						colTotal = 1
					}
					colTotal *= numbers[i][j]
				}
			}
			grandTotal += colTotal
		}
	}
	
	return grandTotal
}
