package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"strings"
	"slices"
	"log"
	"fmt"
)

func main() {
	aoc.Harness(run)
}

func parseInput(fname string) (allowed_ids [][2]int, ids []int) {
	lines := strings.Split(strings.TrimSpace(fname), "\n")
	linebreak := slices.Index(lines, "")
	if linebreak == -1 {
		log.Fatal("invalid input: missing line break between allowed IDs and IDs to check")
	}
	// parse allowed IDs
	for _, line := range lines[:linebreak] {
		var min_id, max_id int
		_, err := fmt.Sscanf(line, "%d-%d", &min_id, &max_id)
		if err != nil {
			log.Fatalf("invalid allowed ID line: %s", line)
		}
		allowed_ids = append(allowed_ids, [2]int{min_id, max_id})
	}
	// parse IDs to check
	for _, line := range lines[linebreak+1:] {
		var id int
		_, err := fmt.Sscanf(line, "%d", &id)
		if err != nil {
			log.Fatalf("invalid ID line: %s", line)
		}
		ids = append(ids, id)
	}
	return allowed_ids, ids
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	allowed_ids, ids := parseInput(input)
	if part2 {
		slices.SortFunc(allowed_ids, func(a, b [2]int) int {
			if a[0] == b[0] {
				return a[1] - b[1]
			}
			return a[0] - b[0]
		})

		for i := 1; i < len(allowed_ids); {
			if allowed_ids[i][0] <= allowed_ids[i-1][1]+1 {
				if allowed_ids[i][1] > allowed_ids[i-1][1] {
					allowed_ids[i-1][1] = allowed_ids[i][1]
				}
				allowed_ids = append(allowed_ids[:i], allowed_ids[i+1:]...)
			} else {
				i++
			}
		}
		
		nb_allowed := 0
		for _, r := range allowed_ids {
			nb_allowed += r[1] - r[0] + 1
		}
		return nb_allowed
	}
	// solve part 1 here
	nb_allowed := 0
	for _, id := range ids {
		allowed := false
		for _, allowed_range := range allowed_ids {
			if id >= allowed_range[0] && id <= allowed_range[1] {
				allowed = true
				break
			}
		}
		if allowed {
			nb_allowed++
		}
	}
	return nb_allowed
}
