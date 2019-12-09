/*Package day7 is
OK, I'm really running out of rope here.
This program is proof I don't properly understand Go yet ;)

Very clean solutions on the reddit thread ;(
*/
package day7

import (
    "aoc"
    "intcode"
)

// return copy of vals with element at "index" omitted
func drop(vals []int64, index int64) []int64 {
    vals = append([]int64(nil), vals...)
    return append(vals[:index], vals[index+1:]...)
}

// non-destructive append
func app(a []int64, b... int64) []int64 {
    a = append([]int64(nil), a...)
    return append(a, b...)
}

// NOTE: stack based iteration would be also nice ;)
func permutations(inputs []int64) (res [][]int64) {
    var worker func(ins, outs []int64, vals *[][]int64)
    worker = func(ins, outs []int64, vals *[][]int64) {
        if len(ins) == 0 {
            *vals = append(*vals, outs)
        }
        for i := range ins {
            val, rest := ins[int64(i)], drop(ins, int64(i))
            worker(rest, app(outs, val), vals)
        }
    }

    worker(inputs, nil, &res)
    return
}

func checkPerm(program, perm []int64) (val int64) {
    // iteratively computing amplifier series
    for _, phase:= range perm {
        state := intcode.Exec(program, []int64{phase, val})
        val = state.OutputVals[0]
    }
    return
}

func checkPermFeedback(program, perm []int64) int64 {
    chans := [](chan int64){}
    for range perm {
        chans = append(chans, make(chan int64, 1))
    }

    done := make(chan bool, len(perm))

    for i, phase := range perm {
        cp := append([]int64(nil), program...)
        j := (i+1) % len(perm)
        state := intcode.MakeState(cp, chans[i], chans[j])

        go state.Run(done)

        chans[i] <- phase
    }

    var initialVal int64 = 0
    chans[0] <- initialVal

    // wait for all states to finish
    for range perm {
        <- done
    }
    close(done)

    // NOTE: intcode closes all output channels
    // since all inputs are also outputs they are already closed

    return <- chans[0]
}


// RunPart runs the program with given inputs
func RunPart(program, ps []int64, f func(a, b []int64) int64) (max int64) {
    for _, perm := range permutations(ps) {
        val := f(program, perm)
        if val > max {
            max = val
        }
    }
    return
}

// Main executes the code for the day 2 exercise
func Main() {
    program := aoc.ReadCommaInts64("data/day7_input.txt")
    phaseSettings := []int64{0, 1, 2, 3, 4}
    aoc.CheckMain("day7.1", int(RunPart(program, phaseSettings, checkPerm)), 46014)

    phaseSettings = []int64{5, 6, 7, 8, 9}
    aoc.CheckMain("day7.2", int(RunPart(program, phaseSettings, checkPermFeedback)),
        19581200)
}
