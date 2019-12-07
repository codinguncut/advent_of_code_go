/*Package day7 is
OK, I'm really running out of rope here.
This program is proof I don't properly understand Go yet ;)

Very clean solutions on the reddit thread ;(
*/
package day7

import (
    "fmt"
    "aoc"
    "intcode"
)

// I had weird artifacts when straight up "append"ing
// I don't understand copy semantics of slices yet
func drop(vals []int, index int) []int {
    cp := make([]int, len(vals))
    copy(cp, vals)
    return append(cp[:index], cp[index+1:]...)
}

// defensive append?!?
// *argh*
func app(a []int, b... int) []int {
    cp := make([]int, len(a))
    copy(cp, a)
    return append(cp, b...)
}

// *argh*
func permutations(inputs []int) [][]int {
    vals := make(chan []int)

    // my clunky attempt at a (recursive) generator
    // FIXME: would be much better with a stack and iteration
    // declare worker beforehand to enable recursion ;)
    var worker func(vals, inputs []int, res chan []int)
    worker = func(vals, inputs []int, res chan []int) {
        // leaf node
        if len(inputs) == 0 {
            res <- vals
            return
        }
        for i := range inputs {
            val := inputs[i]
            rest := drop(inputs, i)
            worker(app(vals, val), rest, res)
        }
        // root node
        if len(vals) == 0 {
            close(res)
        }
    }
    go worker(nil, inputs, vals)

    res := []([]int){}
    for v := range vals {
        res = append(res, v)
    }
    return res
}

func checkPerm(program, perm []int) int {
    val := 0
    // iteratively computing amplifier series
    for _, phase:= range perm {
        state := intcode.Exec(program, []int{phase, val})
        val = state.OutputVals[0]
    }
    return val
}

/*
Creating daisy chain of "State"s.
Starting with a placeholder for first input, and finally
    replacing the placeholder with the last output

NOTE: It would have been a million times easier to stop each Intcode
    whenever it is looking for input instead of using actual concurrency,
    but oh well ;)
*/
func checkPermFeedback(program, perm []int) int {
    inputChannels := [](chan int){}
    for range perm {
        // last output needs to be buffered, since it writes one
        //  additional final value
        ch := make(chan int, 1)
        inputChannels = append(inputChannels, ch)
    }

    finished := make(chan bool)

    for i := range perm {
        programCopy := make([]int, len(program))
        copy(programCopy, program)

        j := (i+1) % len(perm)
        state := intcode.MakeState(programCopy,
            inputChannels[i], inputChannels[j])

        go state.Run(finished)

        // sending initial phase param into each state
        inputChannels[i] <- perm[i]
    }

    // sending initial 0 value into first "State"
    inputChannels[0] <- 0

    // wait for all states to finish
    for range perm {
        <- finished
    }
    close(finished)

    res := <- inputChannels[0]

    // NOTE: intcode closes all output channels
    // since all inputs are also outputs they are already closed

    return res
}


// RunPart runs the program with given inputs
func RunPart(program []int, ps []int, f func([]int, []int) int) int {
    max := 0
    for _, perm := range permutations(ps) {
        val := f(program, perm)
        if val > max {
            max = val
        }
    }
    return max
}

// Main executes the code for the day 2 exercise
func Main() {
    program := aoc.ReadCommaInts("data/day7_input.txt")
    phaseSettings := []int{0, 1, 2, 3, 4}
    fmt.Println("day7.1", RunPart(program, phaseSettings, checkPerm))

    phaseSettings = []int{5, 6, 7, 8, 9}
    fmt.Println("day7.2", RunPart(program, phaseSettings, checkPermFeedback))

    // aoc.CheckMain("day5.1", RunPart(program, []int{1}), 15314507)
    // aoc.CheckMain("day5.2", RunPart(program, []int{5}), 652726)
}
