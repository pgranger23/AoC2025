package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"strings"
	"log"
)

func main() {
	aoc.Harness(run)
}

func parseInput(fname string) (splitters [][]bool, startPos int) {
	for _, line := range strings.Split(strings.TrimSpace(fname), "\n") {
		lineSplit := make([]bool, len(line))
		for i, ch := range line {
			if ch == '^' {
				lineSplit[i] = true
			} else if ch == '.' {
				lineSplit[i] = false
			} else if ch == 'S' {
				startPos = i
			} else {
				log.Fatalf("invalid character in input: %c", ch)
			}
		}
		splitters = append(splitters, lineSplit)
	}

	if startPos == 0 {
		log.Fatal("invalid input: missing start position 'S'")
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
	splitters, startPos := parseInput(input)

	if part2 {
		var pathsToProcess map[[2]int]int = make(map[[2]int]int)
		pathsToProcess[[2]int{0, startPos}] = 1
		totalPaths := 0

		for len(pathsToProcess) > 0 {
			var currentPath [2]int
			var count int
			for key, val := range pathsToProcess {
				currentPath = key
				count = val
				break
			}
			delete(pathsToProcess, currentPath)
			currentLine := currentPath[0]
			currentPos := currentPath[1]


			if currentLine == len(splitters) {
				totalPaths += count
				continue
			}

			if splitters[currentLine][currentPos] {
				if currentPos > 0 {
					pathsToProcess[[2]int{currentLine + 1, currentPos - 1}] += count
				}
				if currentPos < len(splitters[0])-1 {
					pathsToProcess[[2]int{currentLine + 1, currentPos + 1}] += count
				}
			} else {
				pathsToProcess[[2]int{currentLine + 1, currentPos}] += count
			}
		}

		return totalPaths
	}
	
	currentState := make([]bool, len(splitters[0]))
	currentState[startPos] = true

	nbSplits := 0

	for _, line := range splitters {
		nextState := make([]bool, len(currentState))
		for i := 0; i < len(currentState); i++ {
			if currentState[i] {
				if line[i] {
					nbSplits++
					if i > 0 {
						nextState[i-1] = true
						
					}
					if i < len(currentState)-1 {
						nextState[i+1] = true
					}
				} else {
					nextState[i] = true
				}
			}
		}
		currentState = nextState
	}



	return nbSplits
}
