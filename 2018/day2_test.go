package main

import "testing"

func TestSingle(t *testing.T) {
    tables := []struct {
        box_id string
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
        if mp := checksum_single(table.box_id); mp[2] != table.twos || mp[3] != table.threes {
            t.Errorf("result was incorrect, val: %v, twos: %v, want_twos: %v, threed %v, want_threes: %v", table.box_id, mp[2], table.twos, mp[3], table.threes)
        }
    }
}

func TestChecksumBoxes(t *testing.T) {
    boxes := []count_map{
        {},
        {2: true, 3: true},
        {2: true},
        {3: true},
        {2: true},
        {2: true},
        {3: true},
    }
    want := 12
    if got := checksum_boxes(boxes); got != want {
        t.Errorf("result was incorrect, got: %v, want: %v", got, want)
    }
}
