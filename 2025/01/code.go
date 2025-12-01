package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"fmt"
	"strings"
)

func main() {
	aoc.Harness(run)
}

func parseInput(fname string) []int {
	var result []int
	for _, line := range strings.Split(strings.TrimSpace(fname), "\n") {
		// first char of line is a letter, read it
		char := line[0]
		//Rest is an integer, read it
		var number int
		_, err := fmt.Sscanf(line[1:], "%d", &number)
		if err != nil {
			panic(err)
		}
		if char == 'L' {
			result = append(result, -number)
		} else {
			result = append(result, number)
		}
		
	}
	return result
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	values := parseInput(input)
	counter := 50
	nb_crosses := 0
	if part2 {
		for _, v := range values {
			old_counter := counter
			counter += v
			if counter > 0 {
				nb_crosses += counter / 100
			} else if counter < 0 {
				nb_crosses += -counter / 100
				if old_counter != 0 {
					nb_crosses += 1
				}
			} else if counter == 0 {
				nb_crosses += 1
			}
			counter = counter%100
			if counter < 0 {
				counter += 100
			}
			// fmt.Println(counter, nb_crosses)
		}
		return nb_crosses
	}
	// solve part 1 here
	
	for _, v := range values {
		counter += v
		counter = counter%100
		if counter == 0 {
			nb_crosses += 1
		}
	}
	return nb_crosses
}