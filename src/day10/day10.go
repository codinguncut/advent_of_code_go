package day10

import (
    "aoc"
    "fmt"
    "math"
    "sort"
)

type vector struct {
    x, y float64

}

// compare two floats for approximate equality
func floatEqual(a, b float64) bool {
    epsilon := 1e-5
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
    if angle < 0 {
        // convert from [-pi, pi] to [0, 2pi]
        angle = math.Mod(angle + 2.0 * math.Pi, 2.0 * math.Pi)
    }
    return polar{angle, vec.magnitude()}
}

func (vec vector) unit() vector {
    mag := vec.magnitude()
    return vector{vec.x / mag, vec.y / mag}
}

// approximate vector equality
func (vec vector) equal(other vector) bool {
    epsilon := 1e-5
    dx, dy := math.Abs(vec.x - other.x), math.Abs(vec.y - other.y)
    if dx < epsilon && dy < epsilon {
        return true
    }
    return false
}

// approximate polar equality
func (pol polar) equal(other polar) bool {
    epsilon := 1e-5
    da := math.Abs(pol.angle - other.angle)
    dd := math.Abs(pol.dist - other.dist)
    if da < epsilon && dd < epsilon {
        return true
    }
    return false
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

    collides := 0
    for i, a := range deltas {
        for j := i+1; j < len(deltas); j++ {
            b := deltas[j]
            if a.equal(b) {
                // a is collides with b
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

// calculate number and index of asteroid that can see most others
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
    index := 1
    curr, pols := pols[0], pols[1:]
    for {
        if len(pols) == 0 {
            panic("didn't get to target int")
        }

        toDrop := []int{}
        for i, p := range pols {
            // same angle as previous => skip
            if floatEqual(p.angle, curr.angle) {
                continue
            }
            index++
            if index == target {
                return p.vector().add(origin)
            }
            curr = p 
            toDrop = append(toDrop, i)
        }
        // FIXME: can I drop from pols during range iteration?
        for _, i := range toDrop {
            pols = dropPols(pols, i)
        }
    }
}

// Main is
func Main() {
    vecs := readLines(aoc.ReadLines("data/day10_input.txt"))
    maxVal, maxIdx := part1(vecs)
    maxVec := vecs[maxIdx]
    fmt.Println("day10.1", maxVal, maxVec)


    pols := getPolars(maxVec, drop(vecs, maxIdx))
    fmt.Println("day10.2", shootAsteroids(maxVec, pols, 200))
}
