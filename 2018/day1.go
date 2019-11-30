package main

import (
    "fmt"
)

func calc_frequency(nums []int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

// generate next value in sequence
// uses a closure to keep state between calls
func make_next_1() (func(int) int) {
    total := 0
    return func(val int) int {
        total += val
        return total
    }
}

// FIXME: use channels as "make_next_2"

// find recurring element from next_val sequence
func find_recurring(nums []int, next_val func(int) int) int {
    seen := map[int]bool{0: true}
    // NOTE: can loop forever
    //  no trivial way to bound
    for {
        for _, val  := range nums {
            x := next_val(val)
            if seen[x] {
                return x
            }
            seen[x] = true
        }
    }
}

func recurring_1(nums []int) int {
    return find_recurring(nums, make_next_1())
}

// day 1 main
func day1_main() {
    vals := read_file_ints("data/day1_input.txt")
    fmt.Println("day 1.1", calc_frequency(vals))

    fmt.Println("day 1.2", recurring_1(vals))
}
