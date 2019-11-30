package main

import (
    "fmt"
    "strconv"
    "io/ioutil"
    "strings"
)

// generate next value in sequence
// uses a closure to keep state between calls
func make_next_1() (func(int) int) {
    total := 0
    return func(val int) int {
        total += val
        return total
    }
}

// FIXME: use channels as "make_next_2"

// find recurring element from next_val sequence
func find_recurring(nums []int, next_val func(int) int) int {
    seen := map[int]bool{0: true}
    // NOTE: can loop forever
    for {
        for _, val  := range nums {
            x := next_val(val)
            if seen[x] {
                return x
            }
            seen[x] = true
        }
    }
}

func recurring_1(nums []int) int {
    return find_recurring(nums, make_next_1())
}

func calc_frequency(nums []int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

// read newline separated integers from file
func read_file_ints(fname string) []int {
    dat, err := ioutil.ReadFile(fname)
    check(err)

    vals := []int{}
    strs := strings.Split(string(dat), "\n")
    for _, str := range strs {
        if strings.TrimSpace(str) == "" {
            continue
        }
        val, err := strconv.ParseInt(str, 10, 32)
        check(err)
        vals = append(vals, int(val))
    }
    return vals
}

// day 1 main
func day1() {
    vals := read_file_ints("data/day1_input.txt")
    fmt.Println("day 1.1", calc_frequency(vals))

    fmt.Println("day 1.2", recurring_1(vals))
}
