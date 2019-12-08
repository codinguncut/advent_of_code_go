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

// I had weird artifacts when straight up "append"ing
// I don't understand copy semantics of slices yet
func drop(vals []int, index int) []int {
    // FIXME!!!!
    vals = append([]int(nil), vals...)
    return append(vals[:index], vals[index+1:]...)
}

// defensive append?!?
// *argh*
func app(a []int, b... int) []int {
    // FIXME!!!!
    a = append([]int(nil), a...)
    return append(a, b...)
}

func permutations(inputs []int) (res [][]int) {
    var worker func(ins, outs []int, vals *[][]int)
    worker = func(ins, outs []int, vals *[][]int) {
        if len(ins) == 0 {
            *vals = append(*vals, outs)
        }
        for i := range ins {
            val, rest := ins[i], drop(ins, i)
            worker(rest, app(outs, val), vals)
        }
    }

    worker(inputs, []int(nil), &res)
    return
}

func checkPerm(program, perm []int) (val int) {
    // iteratively computing amplifier series
    for _, phase:= range perm {
        state := intcode.Exec(program, []int{phase, val})
        val = state.OutputVals[0]
    }
    return
}

func checkPermFeedback(program, perm []int) int {
    chans := [](chan int){}
    for range perm {
        chans = append(chans, make(chan int, 1))
    }

    done := make(chan bool, len(perm))

    for i, phase := range perm {
        cp := append([]int(nil), program...)
        j := (i+1) % len(perm)
        state := intcode.MakeState(cp, chans[i], chans[j])

        go state.Run(done)

        chans[i] <- phase
    }

    // sending initial 0 value into first "State"
    initialVal := 0
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
func RunPart(program []int, ps []int, f func([]int, []int) int) (max int) {
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
    program := aoc.ReadCommaInts("data/day7_input.txt")
    phaseSettings := []int{0, 1, 2, 3, 4}
    aoc.CheckMain("day7.1", RunPart(program, phaseSettings, checkPerm), 46014)

    phaseSettings = []int{5, 6, 7, 8, 9}
    aoc.CheckMain("day7.2", RunPart(program, phaseSettings, checkPermFeedback),
        19581200)
}
