package main

import (
    "fmt"
)

// Opcode represents operations that the processor supports
type Opcode int

const (
    // OpAdd (from1, from2, target): target = from1 + from2
    OpAdd Opcode = 1
    // OpMul (from1, from2, target): target = from1 * from2
    OpMul Opcode = 2
    // OpDone signifies that a program has terminated
    OpDone Opcode = 99
)

// State data type is contains the program memory and the instruction pointer
type State struct {
    pos []int // memory
    ip int // instruction pointer
}

func (state State) opcode() Opcode {
    return Opcode(state.pos[state.ip])
}

// Eval evaluates opcode at "ip" and mutate "state" accordingly
func (state *State) Eval() (opcodeLength int) {
    ip := state.ip
    from1, from2, target := state.pos[ip+1], state.pos[ip+2], state.pos[ip+3]
    a, b := state.pos[from1], state.pos[from2]
    opLen := 4

    switch state.opcode() {
    case OpAdd:
        state.pos[target] = a+b
        return opLen
    case OpMul:
        state.pos[target] = a*b
        return opLen
    default:
        panic("unknown opcode")  // FIXME: which one
    }
}

// Run runs the program starting at "state" and mutates it
func (state *State) Run() {
    for {
        if state.opcode() == OpDone {
            return
        }
        opLen := state.Eval()
        state.ip += opLen
    }
}

// Exec take a program, creates a State and runs it
func Exec(program []int) *State {
    state := &State{pos: program, ip: 0}
    state.Run()
    return state
}

// RunPart1 takes two initializing values and a program, sets cells 1 and 2,
//  and returns the resulting cell 0 value
func RunPart1(i int, j int, program []int) int {
    // initialize "noun" and "verb" at positions 1 and 2
    program[1] = i
    program[2] = j

    state := Exec(program)
    return state.pos[0]
}

// RunPart2 loads the input from the file and brute-forces all possible
//  values for cells 1 and 2 to find a resulting target value
func RunPart2() int {
    program := read_comma_ints("data/day2_input.txt")

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

// Day2Main executes the code for the day 2 exercise
func Day2Main() {
    program := read_comma_ints("data/day2_input.txt")

    fmt.Println("day2.1", RunPart1(12, 2, program))
    fmt.Println("day2.2", RunPart2())
}
