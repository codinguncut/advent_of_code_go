package day5

import (
    "aoc"
    "intcode"
)

// RunPart runs the program with given inputs
func RunPart(program []int, inputs []int) int {
    state := intcode.Exec(program, inputs)
    outs := state.Outputs
    return outs[len(outs)-1]
}

// Main executes the code for the day 2 exercise
func Main() {
    program := aoc.ReadCommaInts("data/day5_input.txt")

    aoc.CheckMain("day5.1", RunPart(program, []int{1}), 15314507)
    aoc.CheckMain("day5.2", RunPart(program, []int{5}), 652726)
}
