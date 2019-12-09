package aoc

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

// ReadFile reads a file as string
func ReadFile(fname string) string {
    dat, err := ioutil.ReadFile(fname)
    Check(err)
    return string(dat)
}

// ReadLines reads non-empty lines as strings from file
func ReadLines(fname string) []string {
    dat, err := ioutil.ReadFile(fname)
    Check(err)

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

// ReadCommaInts reads comma-separated integers from file
func ReadCommaInts(fname string) []int {
    str := strings.TrimSpace(ReadFile(fname))

    vals := []int{}
    for _, str := range strings.Split(str, ",") {
        val, err := strconv.Atoi(str)
        Check(err)
        vals = append(vals, int(val))
    }
    return vals
}

// ReadCommaInts64 reads comma-separated integers from file
func ReadCommaInts64(fname string) (vals []int64) {
    str := strings.TrimSpace(ReadFile(fname))

    for _, str := range strings.Split(str, ",") {
        val, err := strconv.ParseInt(str, 10, 64)
        Check(err)
        vals = append(vals, val)
    }
    return vals
}


// ReadFileInts reads integers from file, one per line
func ReadFileInts(fname string) []int {
    str := strings.TrimSpace(ReadFile(fname))

    vals := []int{}
    for _, str := range strings.Split(str, "\n") {
        val, err := strconv.Atoi(str)
        Check(err)
        vals = append(vals, int(val))
    }
    return vals
}

// Check converts Errors into panic()
func Check(e error) {
    if e != nil {
        panic(e)
    }
}

// CheckOk panics of ok is not true
func CheckOk(ok bool, str string) {
    if !ok {
        panic(str)
    }
}

// CheckMain provides standardized reporting for each day results
func CheckMain(label string, got int, want int) {
    fmt.Println(label, got)
    CheckOk(got == want, fmt.Sprintf("didn't match, got: %d, want: %d",
        got, want))
}
