package main

import (
    "fmt"
    // "github.com/texttheater/golang-levenshtein/levenshtein"
)

type countMap = map[int]bool

// test whether there are letters in box_id that occur N times
func checksumSingle(boxID string) (res countMap) {
    counts := map[rune]int{}
    for _, char := range boxID {
        counts[char]++
    }

    res = countMap{}
    for _, v := range counts {
        res[v] = true
    }
    return
}

// multiply number of box_id's with 2 letters occuring
//  with number of 3 letters occurring
func checksumBoxes(cmaps []countMap) int {
    twos := 0
    threes := 0
    for _, cmap := range cmaps {
        if cmap[2] {
            twos++
        }
        if cmap[3] {
            threes++
        }
    }
    return twos * threes
}

// StringDeltas computes indices of non-corresponding characters in two strings
func StringDeltas(a, b string) []int {
    deltas := []int{}
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            deltas = append(deltas, i)
            // we don't care about >1 deltas
            if len(deltas) > 1 {
                return nil
            }
        }
    }
    return deltas
}

// Pairwise makes O(n^2) string comparisons to find the string with
//  distance == 1
func Pairwise(strs []string) (common string) {
    for i := range strs {
        for j := i+1; j < len(strs); j++ {
            deltas := StringDeltas(strs[i], strs[j])
            if len(deltas) == 1 {
                index := deltas[0]
                common = strs[i][:index] + strs[i][index+1:]
                return
            }
        }
    }
    panic("didn't find the matching strings")
}

func day2Main() {
    boxIds := read_lines("data/day2_input.txt")
    cmaps := []countMap{}
    for _, bid:= range boxIds {
        cmaps = append(cmaps, checksumSingle(bid))
    }
    checksum := checksumBoxes(cmaps)
    fmt.Println("day2.1: checksum", checksum)

    fmt.Println("day2.2:", Pairwise(boxIds))
}
