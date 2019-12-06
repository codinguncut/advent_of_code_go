package day6

import (
    "aoc"
    "strings"
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
    for _, line := range lines {
        strs := strings.Split(string(line), ")")
        a, b := strs[0], strs[1]
        orbs.parents[b] = a
    }
}

// calculate paths from all nodes to COM
func (orbs *orbits) calcPath(name string) []string {
    parent, ok := orbs.parents[name]
    if (!ok) {
        // we have hit "COM"
        return []string{}
    }
    
    // path hasn't been precomputed yet
    if _, ok := orbs.paths[name]; !ok {
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

    you, ok := orbs.paths["YOU"]
    aoc.CheckOk(ok, "paths[YOU]")

    san, ok := orbs.paths["SAN"]
    aoc.CheckOk(ok, "paths[SAN]")

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

    aoc.CheckMain("day6.1", orbs.calcPathLens(), 402879)
    aoc.CheckMain("day6.2", orbs.part2(), 484)
}
