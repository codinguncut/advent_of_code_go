package main

import (
    "fmt"
)

type count_map = map[int]bool

// test whether there are letters in box_id that occur N times
func checksum_single(box_id string) (res count_map) {
    counts := map[rune]int{}
    for _, char := range box_id {
        counts[char] += 1
    }

    res = count_map{}
    for _, v := range counts {
        res[v] = true
    }
    return
}

// multiply number of box_id's with 2 letters occuring
//  with number of 3 letters occurring
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

func day2_main() {
    box_ids := read_lines("data/day2_input.txt")
    cmaps := []count_map{}
    for _, box_id := range box_ids {
        cmaps = append(cmaps, checksum_single(box_id))
    }
    checksum := checksum_boxes(cmaps)
    fmt.Println("day2.1: checksum", checksum)
}
