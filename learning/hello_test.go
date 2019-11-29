package main

import (
    "testing"
    "sort"
)

func TestSorting(t *testing.T) {
    a := []int{7, 2, 4}
    b := make([]int, len(a))
    copy(a, b)
    sort.Ints(b)

    if len(a) != len(b) {
        t.Error("a and b should have same length")
    }
}

