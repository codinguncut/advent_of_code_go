package main

import (
    "fmt"
    "strings"
    "strconv"
)

// calculate absolute value of "x"
func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

type Coord struct {
    x, y int
}

func (this Coord) Add(other Coord) Coord {
    return Coord{this.x + other.x, this.y + other.y}
}

func (this Coord) CalcManhattan() int {
    return Abs(this.x) + Abs(this.y)
}

///

type PathSegment struct {
    dir Coord
    length int
}

type Path = []PathSegment

var DirLookup = map[string]Coord{
    "R": Coord{ 1,  0},
    "U": Coord{ 0,  1},
    "L": Coord{-1,  0},
    "D": Coord{ 0, -1},
}
var Origin Coord = Coord{0, 0}

///

type Visited map[Coord]int

func ParseSegment(path string) PathSegment {
    length, err := strconv.ParseInt(path[1:], 10, 32)
    check(err)
    return PathSegment{DirLookup[path[0:1]], int(length)} 
}

func ParsePath(line string) Path {
    path := Path{}
    for _, part := range strings.Split(line, ",") {
        path = append(path, ParseSegment(part))
    }
    return path
}

func ReadPaths() (p1, p2 Path) {
    lines := read_lines("data/day3_input.txt")
    p1 = ParsePath(lines[0])
    p2 = ParsePath(lines[1])
    return
}

func PathToVisited(path Path) Visited {
    locs := Visited{}
    curr_loc := Origin
    step := 0
    for _, p := range path {
        for i := 0; i < p.length; i++ {
            step += 1
            curr_loc = curr_loc.Add(p.dir)
            // keep smallest step count
            if locs[curr_loc] == 0 {
                locs[curr_loc] = step
            }
        }
    }
    return locs
}

func IntersectVisited(locs1, locs2 Visited) []Coord {
    coords := []Coord{}
    for k, _ := range locs1 {
        if locs2[k] > 0 {
            coords = append(coords, k)
        }
    }
    return coords
}

func FindClosest(coords []Coord) int {
    min := coords[0].CalcManhattan()
    for _, coord := range coords[1:] {
        if coord.CalcManhattan() < min {
            min = coord.CalcManhattan()
        }
    }
    return min
}

func FindShortest(locs1, locs2 Visited) int {
    dists := []int{}
    for k, _ := range locs1 {
        if locs2[k] > 0 {
            dists = append(dists, locs1[k] + locs2[k])
        }
    }

    // two-pass "min" to avoid arbitrary min-seeding value
    min := dists[0]
    for _, v := range dists[1:] {
        if v < min {
            min = v
        }
    }
    return min
}

func Day3Run(p1, p2 Path) int {
    l1 := PathToVisited(p1)
    l2 := PathToVisited(p2)
    return FindClosest(IntersectVisited(l1, l2))
}

func Day3Main() {
    p1, p2 := ReadPaths()
    fmt.Println("day3 closest", Day3Run(p1, p2))

    l1, l2 := PathToVisited(p1), PathToVisited(p2)
    fmt.Println("day3 shortest", FindShortest(l1, l2))
}
