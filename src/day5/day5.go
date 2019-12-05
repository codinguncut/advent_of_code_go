package day5

import (
    "fmt"
    "aoc"
    "intcode"
)

// RunPart runs the program with given inputs
func RunPart(program []int, inputs []int) []int {
    programCopy := make([]int, len(program))
    copy(programCopy, program)
    state := intcode.Exec(programCopy, inputs)
    return state.Outputs
}

// Main executes the code for the day 2 exercise
func Main() {
    program := aoc.ReadCommaInts("data/day5_input.txt")

    fmt.Println("day5.1", RunPart(program, []int{1}))
    fmt.Println("day5.2", RunPart(program, []int{5}))
}
