package day2

import (
    "aoc"
    "intcode"
)

// RunPart1 takes two initializing values and a program, sets cells 1 and 2,
//  and returns the resulting cell 0 value
func RunPart1(i int, j int, program []int) int {
    // TODO: is this copying the slice or the underlying array??
    // could use `append([]int(nil), program...)`
    memory := make([]int, len(program))
    copy(memory, program)

    // initialize "noun" and "verb" at positions 1 and 2
    memory[1] = i
    memory[2] = j

    state := intcode.Exec(memory, nil)
    return state.Mem[0]
}

// RunPart2 loads the input from the file and brute-forces all possible
//  values for cells 1 and 2 to find a resulting target value
func RunPart2(target int) int {
    program := aoc.ReadCommaInts("data/day2_input.txt")

    for i := range [99]int{} {
        for j := range [99]int{} {
            res := RunPart1(i, j, program)

            // reached target value
            if res == target {
                nounAndVerb := i*100+j
                return nounAndVerb
            }
        }
    }

    panic("no solution found")
}

// Main executes the code for the day 2 exercise
func Main() {
    program := aoc.ReadCommaInts("data/day2_input.txt")

    aoc.CheckMain("day2.1", RunPart1(12, 2, program), 2894520)
    aoc.CheckMain("day2.2", RunPart2(19690720), 9342)
}
