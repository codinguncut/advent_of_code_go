package day10

import (
    "aoc"
    "fmt"
    "math"
    "sort"
)

// smallest acceptable delta for soft float comparison
const epsilon = 1e-5

// TODO: could combine vector and polar...
type vector struct {
    x, y float64

}

// compare two floats for approximate equality
func floatEqual(a, b float64) bool {
    return math.Abs(a - b) < epsilon
}

// polar coordinates
//  angle 0 == upward direction, clockwise
type polar struct {
    angle float64
    dist float64
}

func drop(vals []vector, index int) []vector {
    vals = append([]vector(nil), vals...)
    return append(vals[:index], vals[index+1:]...)
}

func dropPols(vals []polar, index int) []polar{
    vals = append([]polar(nil), vals...)
    return append(vals[:index], vals[index+1:]...)
}

func dropMany(vals []polar, indices map[int]bool) (res []polar) {
    for i, p := range vals {
        if _, ok := indices[i]; !ok {
            res = append(res, p)
        }
    }
    return
}


func (vec vector) magnitude() float64 {
    return math.Sqrt(vec.x*vec.x + vec.y*vec.y)
}

func (vec vector) sub(other vector) vector {
    return vector{vec.x - other.x, vec.y - other.y}
}

func (vec vector) add(other vector) vector {
    return vector{vec.x + other.x, vec.y + other.y}
}

// convert vector to polar coordinates
func (vec vector) polar() polar {
    // NOTE: passing (x, -y) instead of (y, x)
    angle := math.Atan2(vec.x, -vec.y)

    // convert from [-pi, pi] to [0, 2*pi]
    angle = math.Mod(angle + 2.0 * math.Pi, 2.0 * math.Pi)
    return polar{angle, vec.magnitude()}
}

func (vec vector) unit() vector {
    mag := vec.magnitude()
    return vector{vec.x / mag, vec.y / mag}
}

// approximate vector equality
func (vec vector) equal(other vector) bool {
    dx, dy := math.Abs(vec.x - other.x), math.Abs(vec.y - other.y)
    return (dx < epsilon && dy < epsilon)
}

// approximate polar equality
func (pol polar) equal(other polar) bool {
    da := math.Abs(pol.angle - other.angle)
    dd := math.Abs(pol.dist - other.dist)
    return (da < epsilon && dd < epsilon)
}

// convert polar to vector representation
func (pol polar) vector() vector {
    return vector{
        math.Round(math.Sin(pol.angle) * pol.dist),
        math.Round(-math.Cos(pol.angle) * pol.dist),
    }
}

type polarOrder []polar;

func (po polarOrder) Len() int {
    return len(po)
}

func (po polarOrder) Swap(i, j int) {
    po[i], po[j] = po[j], po[i]
}

// compares polar coordinates first by angle then by dist
func (po polarOrder) Less(i, j int) bool {
    if po[i].angle < po[j].angle {
        return true
    }
    if floatEqual(po[i].angle, po[j].angle) {
        return po[i].dist < po[j].dist
    }
    return false
}

// convert string lines to vector slice of asteroid locations
func readLines(lines []string) (vecs []vector) {
    for y, line := range lines {
        for x, char := range line {
            if char == '#' {
                vecs = append(vecs, vector{float64(x), float64(y)})
            }
        }
    }
    return
}

// count asteroids that "vec" can see directly
func countSingle(vec vector, others []vector) (count int) {
    // calculate unit vectors of deltas from "vec" to "others"
    deltas := []vector{}
    for _, o := range others {
        deltas = append(deltas, vec.sub(o).unit())
    }

    // count asteroid "a" that overlap with any other asteroids
    collides := 0
    for i, a := range deltas {
        for j := i+1; j < len(deltas); j++ {
            b := deltas[j]
            if a.equal(b) {
                // a is collides with at least one asteroid
                //  no need to check if a collides with others
                collides++
                break
            }
        }
    }
    return len(others) - collides
}

// count visible counts for all vectors
func countVisible(vecs []vector) (counts []int) {
    counts = make([]int, len(vecs))
    for i, v := range vecs {
        counts[i] = countSingle(v, drop(vecs, i))
    }
    return
}

// calculate number and index of the asteroid that can see most others
func part1(vecs []vector) (max int, maxIdx int) {
    counts := countVisible(vecs)
    max = counts[0]
    maxIdx = 0
    for i, c := range counts {
        if c > max {
            max = c
            maxIdx = i
        }
    }
    return
}

// convert "others" to polar coordinates relative to "origin"
//  and sort first by angle and then dist
func getPolars(origin vector, others []vector) (pols []polar) {
    for _, o := range others {
        pols = append(pols, o.sub(origin).polar())
    }
    sort.Sort(polarOrder(pols))    
    return
}

// shoot asteroids in clockwise fashion, skipping when one asteroid
//  is hidden behind the one being shot
func shootAsteroids(origin vector, pols []polar, target int) vector {
    index := 0
    for {
        curr := polar{math.NaN(), math.NaN()} // impossible value for cmp

        toDrop := map[int]bool{}
        for i, p := range pols {
            // same angle as previous => skip
            if floatEqual(p.angle, curr.angle) {
                continue
            }
            index++
            // fmt.Println(p.vector().add(origin))
            if index == target {
                return p.vector().add(origin)
            }
            curr = p 
            toDrop[i] = true
        }
        if len(pols) == 0 {
            panic("didn't reach target int")
        }
        if len(toDrop) == 0 {
            panic("nothing to drop")
        }
        // remove asteroid that have been shot
        pols = dropMany(pols, toDrop)
    }
}

// Main is
func Main() {
    vecs := readLines(aoc.ReadLines("data/day10_input.txt"))
    maxVal, maxIdx := part1(vecs)
    maxVec := vecs[maxIdx]
    fmt.Println("day10.1", maxVal, maxVec)

    /*
    vecs = readLines([]string{
        ".#....#####...#..",
        "##...##.#####..##",
        "##...#...#.#####.",
        "..#.....#...###..",
        "..#.#.....#....##",
    })
    maxVal, maxIdx = part1(vecs)
    maxVec = vecs[maxIdx]
    */

    pols := getPolars(maxVec, drop(vecs, maxIdx))
    fmt.Println("day10.2", shootAsteroids(maxVec, pols, 200))
}
