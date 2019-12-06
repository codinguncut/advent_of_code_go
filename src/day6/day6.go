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
func makeOrbits() orbits {
    return orbits{
        parents: map[string]string{},
        paths: map[string]([]string){},
   }
}

// read lines of orbit relationships into "orbits"
func (orbs *orbits) read(lines []string) {
    // create root node with empty path
    orbs.paths["COM"] = []string{}

    for _, line := range lines {
        strs := strings.Split(string(line), ")")
        a, b := strs[0], strs[1]
        orbs.parents[b] = a
    }
}

// calculate paths from all nodes to COM
func (orbs *orbits) calcPath(name string) []string {
    parent := orbs.parents[name]
    
    if orbs.paths[name] == nil {
        if parent == "" {
            panic(fmt.Sprintf("something went terribly wrong %s\n%v",
                name, orbs))
        }

        // TODO: need to be careful with appending to shared slice?
        orbs.paths[name] = append(orbs.calcPath(parent), parent)
    }
    return orbs.paths[name]
}

// calculate the total path length of the graph
func (orbs *orbits) calcPathLens() (totalPathLen int) {
    for k := range orbs.parents {
        path := orbs.calcPath(k)
        totalPathLen += len(path)
    }
    return totalPathLen
}

func (orbs *orbits) part2() int {
    // make sure paths are pre-calculated
    orbs.calcPathLens()

    you := orbs.paths["YOU"]
    san := orbs.paths["SAN"]

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
    orbs := makeOrbits()
    orbs.read(lines)

    fmt.Println("day6.1", orbs.calcPathLens())
    fmt.Println("day6.2", orbs.part2())
}
