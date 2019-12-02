package main

import (
    "fmt"
)

type Opcode int

// define known opcodes
const (
    OP_ADD Opcode = 1
    OP_MUL Opcode = 2
    OP_DONE Opcode = 99
)

type State struct {
    pos []int // memory
    ip int // instruction pointer
}

func (state State) opcode() Opcode {
    return Opcode(state.pos[state.ip])
}

// eval opcode at "ip" and mutate "state" accordingly
func (state *State) eval() (opcode_length int) {
    ip := state.ip
    from1, from2, target := state.pos[ip+1], state.pos[ip+2], state.pos[ip+3]
    a, b := state.pos[from1], state.pos[from2]
    op_len := 4

    switch state.opcode() {
    case OP_ADD:
        state.pos[target] = a+b
        return op_len
    case OP_MUL:
        state.pos[target] = a*b
        return op_len
    default:
        panic("unknown opcode")  // FIXME: which one
    }
}

func (state *State) run() {
    for {
        if state.opcode() == OP_DONE {
            return
        }
        op_len := state.eval()
        state.ip += op_len
    }
}

func exec(program []int) *State {
    state := &State{pos: program, ip: 0}
    state.run()
    return state
}

func run_part1(i int, j int, program []int) int {
    // initialize "noun" and "verb" at positions 1 and 2
    program[1] = i
    program[2] = j

    state := exec(program)
    return state.pos[0]
}

func run_part2() int {
    program := read_comma_ints("data/day2_input.txt")

    // NOTE: brute-force ;(
    for i := range [99]int{} {
        for j := range [99]int{} {
            memory := make([]int, len(program))
            copy(memory, program)

            res := run_part1(i, j, memory)

            // reached target value
            if res == 19690720 {
                fmt.Println("noun", i, "verb", j, "res", i*100+j)
                return res
            }
        }
    }

    return 0
}

func day2_main() {
    program := read_comma_ints("data/day2_input.txt")

    fmt.Println("day2.1", run_part1(12, 2, program))
    fmt.Println("day2.2", run_part2())
}
