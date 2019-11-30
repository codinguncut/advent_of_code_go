package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

type count_map = map[int]bool

func checksum_single(box_id string) (res count_map) {
    counts := make(map[rune]int)
    for _, char := range box_id {
        counts[char] += 1
    }

    res = make(count_map)
    for _, v := range counts {
        res[v] = true
    }
    return
}

// calculate total checksum from individual checksums
func checksum_boxes(cmaps []count_map) int {
    twos := 0
    threes := 0
    for _, cmap := range cmaps {
        if cmap[2] {
            twos += 1
        }
        if cmap[3] {
            threes += 1
        }
    }
    return twos * threes
}

func read_file_strings(fname string) []string {
    dat, err := ioutil.ReadFile(fname)
    check(err)

    strs := strings.Split(string(dat), "\n")
    vals := []string{}  // make()?
    for _, str := range strs {
        str = strings.TrimSpace(str)
        if str == "" {
            continue
        }
        vals = append(vals, str)
    }
    return vals
}

func day2_main() {
    box_ids := read_file_strings("data/day2_input.txt")
    cmaps := []count_map{}
    for _, box_id := range box_ids {
        cmaps = append(cmaps, checksum_single(box_id))
    }
    checksum := checksum_boxes(cmaps)
    fmt.Println("checksum", checksum)
}
