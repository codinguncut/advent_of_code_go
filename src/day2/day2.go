package day2

import (
    "fmt"
    "aoc"
    "intcode"
)

// RunPart1 takes two initializing values and a program, sets cells 1 and 2,
//  and returns the resulting cell 0 value
func RunPart1(i int, j int, program []int) int {
    // initialize "noun" and "verb" at positions 1 and 2
    program[1] = i
    program[2] = j

    state := intcode.Exec(program, nil)
    return state.Mem[0]
}

// RunPart2 loads the input from the file and brute-forces all possible
//  values for cells 1 and 2 to find a resulting target value
func RunPart2() int {
    program := aoc.ReadCommaInts("data/day2_input.txt")

    for i := range [99]int{} {
        for j := range [99]int{} {
            memory := make([]int, len(program))
            copy(memory, program)

            res := RunPart1(i, j, memory)

            // reached target value
            if res == 19690720 {
                fmt.Println("noun", i, "verb", j, "res", i*100+j)
                return res
            }
        }
    }

    return 0
}

// Main executes the code for the day 2 exercise
func Main() {
    program := aoc.ReadCommaInts("data/day2_input.txt")

    fmt.Println("day2.1", RunPart1(12, 2, program))
    fmt.Println("day2.2", RunPart2())
}
