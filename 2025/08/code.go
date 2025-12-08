package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"strings"
	"fmt"
	"sort"
)

func main() {
	aoc.Harness(run)
}

func parseInput(fname string) (coords [][3]int) {
	for _, line := range strings.Split(strings.TrimSpace(fname), "\n") {
		var x, y, z int
		_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			panic(fmt.Sprintf("invalid input line: %s", line))
		}
		coords = append(coords, [3]int{x, y, z})
	}
	return coords
}

func distance(a, b [3]int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1]) + (a[2]-b[2])*(a[2]-b[2])
}



// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	var nbIterations int
	if len(input) < 1000 {
		nbIterations = 10
	} else {
		nbIterations = 1000
	}

	coords := parseInput(input)

	networks := make(map[int][]int)
	nodesNetwork := make([]int, len(coords))
	for id, _ := range coords {
		nodesNetwork[id] = id
		networks[id] = []int{ id }
	}

	distanceMap := make(map[[2]int]int)
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			dist := distance(coords[i], coords[j])
			distanceMap[[2]int{i, j}] = dist
		}
	}

	if part2 {
		var minPair [2]int
		for len(networks) != 1 {
			//Get the minimum distance
			minDist := -1
			
			for pair, dist := range distanceMap {
				if minDist == -1 || dist < minDist {
					minDist = dist
					minPair = pair
				}
			}
			delete(distanceMap, minPair)

			//Merge the networks
			net1 := nodesNetwork[minPair[0]]
			net2 := nodesNetwork[minPair[1]]
			if net1 != net2 {
				//Merge net2 into net1
				for _, node := range networks[net2] {
					nodesNetwork[node] = net1
					networks[net1] = append(networks[net1], node)
				}
				delete(networks, net2)
			}
		}
		return coords[minPair[0]][0]*coords[minPair[1]][0]
	}
	

	for it := 0; it < nbIterations; it++ {
		//Get the minimum distance
		minDist := -1
		var minPair [2]int
		for pair, dist := range distanceMap {
			if minDist == -1 || dist < minDist {
				minDist = dist
				minPair = pair
			}
		}
		delete(distanceMap, minPair)

		//Merge the networks
		net1 := nodesNetwork[minPair[0]]
		net2 := nodesNetwork[minPair[1]]
		if net1 != net2 {
			//Merge net2 into net1
			for _, node := range networks[net2] {
				nodesNetwork[node] = net1
				networks[net1] = append(networks[net1], node)
			}
			delete(networks, net2)
		}
	}

	//Multiplying together the size of the 3 largest networks
	sizes := []int{}
	for _, nodes := range networks {
		sizes = append(sizes, len(nodes))
	}
	// fmt.Println("Network sizes:", sizes)
	//Sorting sizes descending
	sort.Ints(sizes)
	return sizes[len(sizes)-3] * sizes[len(sizes)-2] * sizes[len(sizes)-1]
}