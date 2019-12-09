package day9

import (
    "fmt"
    "aoc"
    "intcode"
)

// RunPart runs the program with given inputs
func RunPart(program []int64, inputs []int64) int64 {
    state := intcode.Exec(program, inputs)
    outs := state.OutputVals
    fmt.Println("opcodes", outs)
    return outs[len(outs)-1]
}

// Main executes the code for the day 2 exercise
func Main() {
    program := aoc.ReadCommaInts64("data/day9_input.txt")

    fmt.Println("day9.1", RunPart(program, []int64{1}))
    fmt.Println("day9.2", RunPart(program, []int64{2}))
}
