package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"strings"
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
			if line[i] == '@' {
				line_digits[i] = 1
			} else {
				line_digits[i] = 0
			}
		}
		result = append(result, line_digits)
	}
	return result
}

var neighbors = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1},          {0, 1},
	{1, -1},  {1, 0},  {1, 1},
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	data := parseInput(input)
	nb_valid := 0

	if part2 {
		stable := false
		for !stable {
			stable = true
			for i := 0; i < len(data); i++ {
				for j := 0; j < len(data[i]); j++ {
					if data[i][j] != 1 {
						continue
					}
					nb_neighbors := 0
					for _, neighbor := range neighbors {
						if i+neighbor[0] < 0 || i+neighbor[0] >= len(data) {
							continue
						}
						if j+neighbor[1] < 0 || j+neighbor[1] >= len(data[i]) {
							continue
						}
						if data[i+neighbor[0]][j+neighbor[1]] == 1 {
							nb_neighbors++
						}
					}
					if nb_neighbors < 4 {
						nb_valid++
						stable = false
						data[i][j] = 0
					}
				}
			}
		}
		return nb_valid
	}
	// solve part 1 here
	
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] != 1 {
				continue
			}
			nb_neighbors := 0
			for _, neighbor := range neighbors {
				if i+neighbor[0] < 0 || i+neighbor[0] >= len(data) {
					continue
				}
				if j+neighbor[1] < 0 || j+neighbor[1] >= len(data[i]) {
					continue
				}
				if data[i+neighbor[0]][j+neighbor[1]] == 1 {
					nb_neighbors++
				}
			}
			if nb_neighbors < 4 {
				nb_valid++
			}
		}
	}
	return nb_valid
}
