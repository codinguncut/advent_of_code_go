package main

import (
    "io/ioutil"
    "strings"
    "strconv"
)

// read non-empty lines from file
func read_lines(fname string) []string {
    dat, err := ioutil.ReadFile(fname)
    check(err)

    strs := strings.Split(string(dat), "\n")
    vals := []string{}
    for _, str := range strs {
        str = strings.TrimSpace(str)
        if str == "" {
            continue
        }
        vals = append(vals, str)
    }
    return vals
}

// read integers from file, one per line
func read_file_ints(fname string) []int {
    vals := []int{}
    for _, str := range read_lines(fname) {
        val, err := strconv.ParseInt(str, 10, 32)
        check(err)
        vals = append(vals, int(val))
    }
    return vals
}

// error handling for lazy people
func check(e error) {
    if e != nil {
        panic(e)
    }
}
