package day6

import (
    "aoc"
    "strings"
    "fmt"
)

// orbits contains orbital relationships as well as cached calculated paths
//  from all objects to COM
type orbits struct {
    parents map[string]string
    paths map[string]([]string)
}

// constructor to create empty maps
func makeorbits() orbits {
    return orbits{
        parents: map[string]string{},
        paths: map[string]([]string){},
   }
}

// read lines of orbit relationships into "orbits"
func (orbits *orbits) read(lines []string) {
    // create root node with empty path
    orbits.paths["COM"] = []string{}

    for _, line := range lines {
        strs := strings.Split(string(line), ")")
        a, b := strs[0], strs[1]
        orbits.parents[b] = a
    }
}

// calculate paths from all nodes to COM
func (orbits *orbits) calcPath(name string) []string {
    parent := orbits.parents[name]
    
    if orbits.paths[name] == nil {
        // TODO: careful with appending shared slice
        orbits.paths[name] = append(orbits.calcPath(parent), parent)
    }
    return orbits.paths[name]
}

// calculate the total path length of the graph
func (orbits *orbits) calcPathLens() (totalPathLen int) {
    for k := range orbits.parents {
        path := orbits.calcPath(k)
        totalPathLen += len(path)
    }
    return totalPathLen
}

func (orbits *orbits) part2() int {
    // make sure paths are pre-calculated
    orbits.calcPathLens()

    you := orbits.paths["YOU"]
    san := orbits.paths["SAN"]

    lenCommon := 0
    for i := 0; i < len(you); i++ {
        if you[i] != san[i] {
            break
        }
        lenCommon++
    }
    
    youTxs := you[lenCommon:]
    sanTxs := san[lenCommon:]

    return len(youTxs) + len(sanTxs)
}

// Main is the entrypoint for the day 6 solutions
func Main() {
    lines := aoc.ReadLines("data/day6_input.txt")
    orbits := makeorbits()
    orbits.read(lines)

    fmt.Println("day6.1", orbits.calcPathLens())
    fmt.Println("day6.2", orbits.part2())
}
