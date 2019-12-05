package intcode

import (
    "fmt"
)

// Opcode represents operations that the processor supports
type Opcode int

// FIXME
// type MemoryIndex int

// opcodes supported by the processor
const (
    OpAdd Opcode = 1
    OpMul Opcode = 2
    OpInput Opcode = 3
    OpOutput Opcode = 4
    OpJumpTrue Opcode = 5
    OpJumpFalse Opcode = 6
    OpLessThan Opcode = 7
    OpEquals Opcode = 8
    OpDone Opcode = 99
)

// ParamMode specifies how parameter values should be interpreted
type ParamMode int

// parameter modes
const (
    ParamPos ParamMode = 0
    ParamImm ParamMode = 1
)

// State data type contains the program memory and the instruction pointer.
//  In addition it manages inputs and outputs of the program
type State struct {
    Mem []int // memory
    IP int // instruction pointer
    Inputs []int
    Outputs []int
    ReadInt (func(*State) int)
    WriteInt (func(*State, int))
}

// MakeState is a constructor that creates a state with non-interactive
//  read and write capability
func MakeState(mem []int, inputs []int) *State {
    readInt := func(state *State) int {
        val := state.Inputs[0]
        state.Inputs = state.Inputs[1:]
        return val
    }
    writeInt := func(state *State, val int) {
        state.Outputs = append(state.Outputs, val)
    }
    state := &State{
        Mem: mem,
        IP: 0,
        Inputs: inputs,
        Outputs: []int{},
        ReadInt: readInt,
        WriteInt: writeInt,
    }
    return state
}

// get current opcode pointed at by instruction pointer
func (state State) opcode() Opcode {
    return Opcode(state.Mem[state.IP])
}


// EvalParam evaluates value at memIndex based on ParamMode
func (state State) EvalParam(memIndex int, mode ParamMode) int {
    value := state.Mem[memIndex]
    switch mode {
    case ParamPos:
        return state.Mem[value]
    case ParamImm:
        return value
    default:
        panic(fmt.Sprintf("unknown ParamMode %v", mode))
    }
}

// BinaryOpVal calculates
func (state *State) BinaryOpVal(ip int, params []ParamMode, f func(int, int) int) {
    a, b := state.EvalParam(ip+1, params[0]), state.EvalParam(ip+2, params[1])
    target := state.Mem[ip+3]
    state.Mem[target] = f(a, b)
    state.IP += 4
}

// ParseOpcode parses raw opcode into opcode + paramModes
func ParseOpcode(raw Opcode) (opcode Opcode, params []ParamMode) {
    params = make([]ParamMode, 3)
    opcode = raw % 100
    op := raw / 100
    for i := range params {
        params[i] = ParamMode(op % 10)
        op /= 10
    }
    return
}

// Eval evaluates opcode at "ip" and mutate "state" accordingly
func (state *State) Eval() {
    ip := state.IP
    opcode, params := ParseOpcode(state.opcode())

    switch opcode {
    case OpAdd:
        state.BinaryOpVal(ip, params, func(a, b int) int {
            return a+b
        })

    case OpMul:
        state.BinaryOpVal(ip, params, func(a, b int) int {
            return a*b
        })

    case OpInput:
        val := state.ReadInt(state)
        target := state.Mem[ip+1]
        state.Mem[target] = val
        state.IP += 2

    case OpOutput:
        val := state.EvalParam(ip+1, params[0])
        state.WriteInt(state, val) 
        state.IP += 2

    case OpJumpTrue:
        a, b := state.EvalParam(ip+1, params[0]), state.EvalParam(ip+2, params[1])
        if a != 0 {
            state.IP = b
        } else {
            state.IP += 3
        }

    case OpJumpFalse:
        a, b := state.EvalParam(ip+1, params[0]), state.EvalParam(ip+2, params[1])
        if a == 0 {
            state.IP = b
        } else {
            state.IP += 3
        }

    case OpLessThan:
        state.BinaryOpVal(ip, params, func(a, b int) int {
            if a < b {
                return 1
            }
            return 0
        })

    case OpEquals:
        state.BinaryOpVal(ip, params, func(a, b int) int {
            if a == b {
                return 1
            }
            return 0
        })

    default:
        panic(fmt.Sprintf("unknown opcode %v", opcode))
    }
}

// Run runs the program starting at "state" and mutates it
func (state *State) Run() {
    for {
        if state.opcode() == OpDone {
            return
        }
        state.Eval()
    }
}

// Exec take a program, creates a State and runs it
func Exec(program []int, inputs []int) *State {
    if inputs == nil {
        inputs = []int{}
    }
    state := MakeState(program, inputs)
    state.Run()
    return state
}
