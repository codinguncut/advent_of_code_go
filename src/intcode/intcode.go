package intcode

import (
    "fmt"
)

// Opcode represents operations that the processor supports
type Opcode int

// MemIdx is used to access values in memory
type MemIdx int

// opcodes supported by the processor
const (
    OpAdd = 1
    OpMul = 2
    OpInput = 3
    OpOutput = 4
    OpJumpTrue = 5
    OpJumpFalse = 6
    OpLessThan = 7
    OpEquals = 8
    OpDone = 99
)

// ParamMode specifies how parameter values should be interpreted
type ParamMode int

// parameter modes
const (
    ParamPos = 0
    ParamImm = 1
)

// State data type contains the program memory and the instruction pointer.
//  In addition it manages inputs and outputs of the program
type State struct {
    Mem map[MemIdx]int
    IP MemIdx
    Inputs <-chan int
    Outputs chan<- int
    OutputVals []int
}

// MakeState is a constructor that creates a state with non-interactive
//  read and write capability
func MakeState(mem []int, inputs, outputs chan int) *State {
    mp := map[MemIdx]int{}
    for i, v := range mem {
        mp[MemIdx(i)] = v
    }
    state := &State{
        Mem: mp,
        IP: 0,
        Inputs: inputs,
        Outputs: outputs,
    }
    return state
}

func (state State) memoryArray() []int {
    max := 0
    for k := range state.Mem {
        if int(k) > max {
            max = int(k)
        }
    }
    arr := make([]int, max+1)
    for k, v := range state.Mem {
        arr[k] = v
    }
    return arr
}

// get current opcode pointed at by instruction pointer
func (state State) opcode() Opcode {
    return Opcode(state.Mem[state.IP])
}


// EvalParam evaluates value at memIndex based on ParamMode
func (state State) EvalParam(memIndex MemIdx, mode ParamMode) int {
    value := state.Mem[memIndex]
    switch mode {
    case ParamPos:
        return state.Mem[MemIdx(value)]
    case ParamImm:
        return value
    default:
        panic(fmt.Sprintf("unknown ParamMode %v", mode))
    }
}

// BinaryOpVal calculates
func (state *State) BinaryOpVal(ip MemIdx, params []ParamMode, f func(int, int) int) {
    a, b := state.EvalParam(ip+1, params[0]), state.EvalParam(ip+2, params[1])
    target := state.Mem[ip+3]
    state.Mem[MemIdx(target)] = f(a, b)
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
        val := <- state.Inputs
        target := state.Mem[ip+1]
        state.Mem[MemIdx(target)] = val
        state.IP += 2

    case OpOutput:
        val := state.EvalParam(ip+1, params[0])
        state.Outputs <- val
        state.IP += 2

    case OpJumpTrue:
        a, b := state.EvalParam(ip+1, params[0]), state.EvalParam(ip+2, params[1])
        if a != 0 {
            state.IP = MemIdx(b)
        } else {
            state.IP += 3
        }

    case OpJumpFalse:
        a, b := state.EvalParam(ip+1, params[0]), state.EvalParam(ip+2, params[1])
        if a == 0 {
            state.IP = MemIdx(b)
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
func (state *State) Run(done chan bool) {
    for {
        if state.opcode() == OpDone {
            // NOTE: this is a bit tricky sequencing-wise, and probably
            //  requires either done or Outputs to be buffered
            done <- true
            close(state.Outputs)
            return
        }
        state.Eval()
    }
}

// Exec take a program, creates a State and runs it
func Exec(program []int, inputVals []int) *State {
    programCopy := make([]int, len(program))
    copy(programCopy, program)

    inputs, outputs, finished := make(chan int), make(chan int), make(chan bool, 1)

    go func() {
        for _, v := range inputVals {
            inputs <- v
        }
        close(inputs)
    }()

    state := MakeState(programCopy, inputs, outputs)
    go state.Run(finished)
    for v := range outputs {
        state.OutputVals = append(state.OutputVals, v)
    }
    close(finished)
    return state
}
