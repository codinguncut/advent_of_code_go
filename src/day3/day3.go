package day3

import (
    "fmt"
    "strings"
    "strconv"
    "aoc"
)

// Abs calculates an absolute value
func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

// Coord represents a coordinate
type Coord struct {
    x, y int
}

// Add adds two coordinates and returns result as a new coordinate 
func (coord Coord) Add(other Coord) Coord {
    return Coord{coord.x + other.x,
                 coord.y + other.y}
}

// CalcManhattan calculates the manhattan distance from origin to a coordinate
func (coord Coord) CalcManhattan() int {
    return (Abs(coord.x) +
            Abs(coord.y))
}

///

// PathSegment represents consecutive steps in the same direction
type PathSegment struct {
    dir Coord
    length int
}

// Path is a series of segments from Origin
type Path = []PathSegment

// DirLookup is used to convert segment letter encoding into coordinates
var DirLookup = map[string]Coord{
    "R": Coord{ 1,  0},
    "U": Coord{ 0,  1},
    "L": Coord{-1,  0},
    "D": Coord{ 0, -1},
}

// Origin represents the (0, 0) point of the coordinate system
var Origin Coord = Coord{0, 0}

///

// Visited contains for each coordinate whether it has been visited by
//  a path and the distance from origin
type Visited map[Coord]int

// ParseSegment converts i.e. "U20" to `PathSegment{dir: up, length: 20}`
func ParseSegment(path string) PathSegment {
    length, err := strconv.Atoi(path[1:])
    aoc.Check(err)
    dir, ok := DirLookup[path[0:1]]
    aoc.CheckOk(ok, fmt.Sprintf("DirLookup %v", path[0:1]))
    return PathSegment{dir, int(length)} 
}

// ParsePath parses a string into its corresponding Path
func ParsePath(line string) (path Path) {
    for _, part := range strings.Split(line, ",") {
        path = append(path, ParseSegment(part))
    }
    return
}

// ReadPaths reads an input file and returns two paths
func ReadPaths() (p1, p2 Path) {
    lines := aoc.ReadLines("data/day3_input.txt")
    return ParsePath(lines[0]), ParsePath(lines[1])
}

// CalcVisited populates the Visited data type from a given Path
func CalcVisited(path Path) (visited Visited) {
    visited = Visited{}
    currLoc := Origin
    step := 0
    for _, segment := range path {
        for i := 0; i < segment.length; i++ {
            step++
            currLoc = currLoc.Add(segment.dir)
            // only keep smallest/ first step count in visited
            if _, ok := visited[currLoc]; !ok {
                visited[currLoc] = step
            }
        }
    }
    return
}

// IntersectVisited finds all coordinates for which the two given
//  Paths/ Visited's overlap
func IntersectVisited(locs1, locs2 Visited) (coords []Coord) {
    for k := range locs1 {
        if _, ok := locs2[k]; ok {
            coords = append(coords, k)
        }
    }
    return
}

// FindClosest finds the closest intersection coordinate by manhattan distance
func FindClosest(coords []Coord) (min int) {
    min = coords[0].CalcManhattan()
    for _, coord := range coords[1:] {
        if coord.CalcManhattan() < min {
            min = coord.CalcManhattan()
        }
    }
    return
}

// FindShortest finds the shortest intersection by total circuit length
func FindShortest(locs1, locs2 Visited) (min int) {
    dists := []int{}
    for k := range locs1 {
        if val, ok := locs2[k]; ok {
            dists = append(dists, locs1[k] + val)
        }
    }

    // two-pass "min" to avoid arbitrary min-seeding value
    min = dists[0]
    for _, v := range dists[1:] {
        if v < min {
            min = v
        }
    }
    return
}

// Run executes part 1 of the day 3 exercise
func Run(p1, p2 Path) int {
    l1, l2 := CalcVisited(p1), CalcVisited(p2)
    return FindClosest(IntersectVisited(l1, l2))
}

// Main executes the code for the day 3 exercise
func Main() {
    p1, p2 := ReadPaths()
    aoc.CheckMain("day3.1", Run(p1, p2), 446)

    l1, l2 := CalcVisited(p1), CalcVisited(p2)
    aoc.CheckMain("day3.2", FindShortest(l1, l2), 9006)
}
