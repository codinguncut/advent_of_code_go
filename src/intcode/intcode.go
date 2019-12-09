package intcode

import (
    "fmt"
)

// Opcode represents operations that the processor supports
type Opcode int

// CellType is an alias type for memory cells
//  NOTE: would have preferred a "new type", but too painful for array casting
type CellType = int64

// MemIdx is used to access values in memory
type MemIdx CellType

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
    OpAdjustRel = 9
    OpDone = 99
)

// ParamMode specifies how parameter values should be interpreted
type ParamMode int

// parameter modes
const (
    ParamPos = 0
    ParamImm = 1
    ParamRel = 2
)

// State data type contains the program memory and the instruction pointer.
//  In addition it manages inputs and outputs of the program
type State struct {
    Mem map[MemIdx]CellType
    IP MemIdx
    RelBase MemIdx
    Inputs <-chan CellType
    Outputs chan<- CellType
    OutputVals []CellType
}

// MakeState is a constructor that creates a state with non-interactive
//  read and write capability
func MakeState(mem []CellType, inputs, outputs chan CellType) *State {
    mp := map[MemIdx]CellType{}
    for i, v := range mem {
        mp[MemIdx(i)] = CellType(v)
    }
    state := &State{
        Mem: mp,
        IP: 0,
        RelBase: 0,
        Inputs: inputs,
        Outputs: outputs,
    }
    return state
}

// get memory content as an array for testing purposes
// TODO: move to intcode_test.go?
func (state State) memoryArray() []CellType {
    // find largest map key
    var max MemIdx = 0
    for k := range state.Mem {
        if k > max {
            max = k
        }
    }

    // allocate array and copy elements
    arr := make([]CellType, max+1)
    for k, v := range state.Mem {
        arr[k] = CellType(v)
    }
    return arr
}

// get current opcode pointed at by instruction pointer
func (state State) opcode() Opcode {
    return Opcode(state.Mem[state.IP])
}


// EvalParam evaluates value at memIndex based on ParamMode
func (state State) EvalParam(memIndex MemIdx, mode ParamMode) CellType {
    value := state.Mem[memIndex]
    switch mode {
    case ParamPos:
        return state.Mem[MemIdx(value)]
    case ParamImm:
        return value
    case ParamRel:
        return state.Mem[state.RelBase + MemIdx(value)]
    default:
        panic(fmt.Sprintf("unknown ParamMode %v", mode))
    }
}

// FIXME: remove duplication with EvalParam
func (state *State) setTarget (memIndex MemIdx, mode ParamMode, value CellType) {
    target := MemIdx(state.Mem[memIndex])
    switch mode {
    case ParamPos:
        state.Mem[target] = value
    case ParamRel:
        state.Mem[state.RelBase + target] = value
    default:
        panic(fmt.Sprintf("unsupported ParamMode %v", mode))
    }
}

// BinaryOpVal calculates
func (state *State) BinaryOpVal(ip MemIdx, params []ParamMode,
        f func(a, b CellType) CellType) {
    a, b := state.EvalParam(ip+1, params[0]), state.EvalParam(ip+2, params[1])
    state.setTarget(ip+3, params[2], f(a, b))
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

// Eval evaluates opcode at "ip" and mutates "state" accordingly
func (state *State) Eval() {
    ip := state.IP
    opcode, params := ParseOpcode(state.opcode())

    switch opcode {
    case OpAdd:
        state.BinaryOpVal(ip, params, func(a, b CellType) CellType {
            return a+b
        })

    case OpMul:
        state.BinaryOpVal(ip, params, func(a, b CellType) CellType{
            return a*b
        })

    case OpInput:
        val := <- state.Inputs
        state.setTarget(ip+1, params[0], val)
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
        state.BinaryOpVal(ip, params, func(a, b CellType) CellType {
            if a < b {
                return 1
            }
            return 0
        })

    case OpEquals:
        state.BinaryOpVal(ip, params, func(a, b CellType) CellType {
            if a == b {
                return 1
            }
            return 0
        })

    case OpAdjustRel:
        a := state.EvalParam(ip+1, params[0])
        state.RelBase += MemIdx(a)
        state.IP += 2

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
func Exec(program []CellType, inputVals []CellType) *State {
    program = append([]CellType(nil), program...)

    inputs, outputs, done:= make(chan CellType), make(chan CellType), make(chan bool, 1)

    // NOTE: current channel design does not allow for interactive inputs
    //  for lack of way to request an input
    go func() {
        for _, v := range inputVals {
            inputs <- v
        }
        close(inputs)
    }()

    state := MakeState(program, inputs, outputs)
    go state.Run(done)
    for v := range outputs {
        state.OutputVals = append(state.OutputVals, CellType(v))
    }
    close(done)
    return state
}
