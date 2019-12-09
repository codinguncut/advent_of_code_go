package day5

import (
    "aoc"
    "intcode"
)

// RunPart runs the program with given inputs
func RunPart(program []int64, inputs []int64) int64 {
    state := intcode.Exec(program, inputs)
    outs := state.OutputVals
    return outs[len(outs)-1]
}

// Main executes the code for the day 2 exercise
func Main() {
    program := aoc.ReadCommaInts64("data/day5_input.txt")

    aoc.CheckMain("day5.1", int(RunPart(program, []int64{1})), 15314507)
    aoc.CheckMain("day5.2", int(RunPart(program, []int64{5})), 652726)
}
