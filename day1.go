package main

import (
    "fmt"
    "math"
)

// CalcFuelRequired calculates fuel requirement per module mass
func CalcFuelRequired(mass int, doRecurse bool) int {
    fuel := int(math.Floor(float64(mass) / 3)) - 2
    if !doRecurse {
        return fuel
    }
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

// Day1Main is the main function to be called for day 1 exercise
func Day1Main() {
    masses := read_file_ints("data/day1_input.txt")
    fmt.Println("day1.1", CalcTotalFuel(masses, false))
    fmt.Println("day1.2", CalcTotalFuel(masses, true))
}
