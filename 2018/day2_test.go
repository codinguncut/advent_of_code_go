package main

import (
    "testing"
    "reflect"
)

func TestSingle(t *testing.T) {
    tables := []struct {
        boxID string
        twos bool
        threes bool
    }{
        {"abcdef", false, false},
        {"bababc", true, true},
        {"abbcde", true, false},
        {"abcccd", false, true},
        {"aabcdd", true, false},
        {"abcdee", true, false},
        {"ababab", false, true},
    }

    for _, table := range tables {
        if mp := checksumSingle(table.boxID); mp[2] != table.twos || mp[3] != table.threes {
            t.Errorf("result was incorrect, val: %v, twos: %v, want_twos: %v, threed %v, want_threes: %v", table.boxID, mp[2], table.twos, mp[3], table.threes)
        }
    }
}

func TestChecksumBoxes(t *testing.T) {
    boxes := []countMap{
        {},
        {2: true, 3: true},
        {2: true},
        {3: true},
        {2: true},
        {2: true},
        {3: true},
    }
    want := 12
    if got := checksumBoxes(boxes); got != want {
        t.Errorf("result was incorrect, got: %v, want: %v", got, want)
    }
}

func TestStringDeltas(t *testing.T) {
    tables := []struct {
        a string
        b string
        deltas []int
    }{
        {"abcde", "abdde", []int{2}},
        {"xxxxx", "yyyyy", nil},
        {"abcde", "abcde", []int{}},
    }
    for _, table := range tables {
        got := StringDeltas(table.a, table.b)
        if !reflect.DeepEqual(got, table.deltas) {
            t.Errorf("result was incorrect, got: %v, want: %v", got, table.deltas)
        }
    }
}
