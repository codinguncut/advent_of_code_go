package day1

import (
    "math"
    "aoc"
)

// CalcFuelRequired calculates fuel requirement per module mass
func CalcFuelRequired(mass int, doRecurse bool) int {
    fuel := int(math.Floor(float64(mass) / 3)) - 2
    if !doRecurse {
        return fuel
    }
    // perform recursive calculation
    if fuel > 0 {
        return fuel + CalcFuelRequired(fuel, true)
    }
    return 0
}

// CalcTotalFuel iterates of module masses and return fuel required
func CalcTotalFuel(masses []int, doRecurse bool) (total int) {
    for _, mass := range masses {
        total += CalcFuelRequired(mass, doRecurse)
    }
    return
}

// Main is the main function to be called for day 1 exercise
func Main() {
    masses := aoc.ReadFileInts("data/day1_input.txt")
    aoc.CheckMain("day1.1", CalcTotalFuel(masses, false), 3365459)
    aoc.CheckMain("day1.2", CalcTotalFuel(masses, true), 5045301)
}
