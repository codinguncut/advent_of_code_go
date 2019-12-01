package main

import (
    "fmt"
    "math"
)

// calculate fuel requirement per module
func fuel_required(mass int) int {
    return int(math.Floor(float64(mass) / 3)) - 2
}

// recursively calculate fuel required including for mass of fuel itself
func fuel_required_2(mass int) int {
    fuel := fuel_required(mass)
    if fuel > 0 {
        return fuel + fuel_required_2(fuel)
    } else {
        return 0
    }
}

func calc_total_fuel(masses []int) (total int) {
    for _, mass := range masses {
        total += fuel_required(mass)
    }
    return
}

func calc_total_fuel_2(masses []int) (total int) {
    for _, mass := range masses {
        total += fuel_required_2(mass)
    }
    return
}

func day1_main() {
    masses := read_file_ints("data/day1_input.txt")
    fmt.Println("day1.1", calc_total_fuel(masses))
    fmt.Println("day1.2", calc_total_fuel_2(masses))
}
