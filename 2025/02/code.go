package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"fmt"
	"strings"
	"math"
)

func main() {
	aoc.Harness(run)
}

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func isInvalid(value_digits string, order int) bool {
	n_digits := len(value_digits)
	if n_digits%order != 0 {
		return false
	}

	chunk_size := n_digits / order
	first_chunk := value_digits[0:chunk_size]
	for i := 1; i < order; i++ {
		if value_digits[i*chunk_size:(i+1)*chunk_size] != first_chunk {
			return false
		}
	}
	return true
}

type IDRange struct {
	min int
	max int
}

func parseInput(fname string) []IDRange {
	var result []IDRange
	for _, line := range strings.Split(strings.TrimSpace(fname), ",") {
		var min, max int
		_, err := fmt.Sscanf(line, "%d-%d", &min, &max)	
		if err != nil {
			panic(err)
		}
		result = append(result, IDRange{min, max})
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
	if part2 {
		sum := 0
		ranges := parseInput(input)
		for _, r := range ranges {
			for i := r.min; i <= r.max; i++ {
				n_digits := int(math.Floor(math.Log10(math.Abs(float64(i))))) + 1
				for order := 2; order <= n_digits; order++ {
					if isInvalid(fmt.Sprintf("%d", i), order) {
						sum += i
						break
					}
				}
			}
		}
		return sum
	}
	sum := 0
	ranges := parseInput(input)
	for _, r := range ranges {
		for i := r.min; i <= r.max; i++ {
			if isInvalid(fmt.Sprintf("%d", i), 2) {
				sum += i
			}
		}
	}
	return sum
}
