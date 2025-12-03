package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"strings"
	"strconv"
	"slices"
	"math"
)

func main() {
	aoc.Harness(run)
}

func parseInput(fname string) [][]int {
	var result [][]int
	for _, line := range strings.Split(strings.TrimSpace(fname), "\n") {
		nb_char := len(line)
		line_digits := make([]int, nb_char)
		for i := 0; i < nb_char; i++ {
			number, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}
			line_digits[i] = number
		}
		result = append(result, line_digits)
	}
	return result
}

func getBiggest(digits []int, nb int) int {
	biggest := slices.Max(digits[0:len(digits)- nb + 1])
	biggestIndex := slices.Index(digits, biggest)
	combined := 0
	combined += biggest * int(math.Pow(10, float64(nb-1)))
	for i:= 1; i < nb; i++ {
		biggest = slices.Max(digits[biggestIndex+1:len(digits)- nb + i + 1])
		biggestIndex = slices.Index(digits[biggestIndex+1:len(digits)- nb + i + 1], biggest) + biggestIndex + 1
		combined += biggest * int(math.Pow(10, float64(nb-i-1)))
	}
	return combined
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	data := parseInput(input)
	sum := 0

	if part2 {
		for _, digits := range data {
			biggest := getBiggest(digits, 12)
			sum += biggest
		}
		return sum
	}
	// solve part 1 here
	
	for _, digits := range data {
		biggest := getBiggest(digits, 2)
		sum += biggest
	}
	return sum
}
