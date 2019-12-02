package main

import (
    "io/ioutil"
    "strings"
    "strconv"
)

func read_file(fname string) string {
    dat, err := ioutil.ReadFile(fname)
    check(err)
    return string(dat)
}

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

// read comma-separated integers from file
func read_comma_ints(fname string) []int {
    str := strings.TrimSpace(read_file(fname))

    vals := []int{}
    for _, str := range strings.Split(str, ",") {
        val, err := strconv.ParseInt(str, 10, 32)
        check(err)
        vals = append(vals, int(val))
    }
    return vals
}

// read integers from file, one per line
func read_file_ints(fname string) []int {
    str := strings.TrimSpace(read_file(fname))

    vals := []int{}
    for _, str := range strings.Split(str, "\n") {
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
