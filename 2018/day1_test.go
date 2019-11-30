package main

import "testing"

func TestCalcFrequency(t *testing.T) {
    tables := []struct {
		vals []int
		want int
	}{
        {[]int{+1, +1, +1}, 3},
        {[]int{+1, +1, -2}, 0},
        {[]int{-1, -2, -3}, -6},
	}

    for _, table := range tables {
        if got := calc_frequency(table.vals); got != table.want {
            t.Errorf("result was incorrect, got: %d, want: %d", got, table.want)
        }
    }
}

func TestFindRecurring(t *testing.T) {
    tables := []struct {
		vals []int
		want int
	}{
        {[]int{+1, -1}, 0},
        {[]int{+3, +3, +4, -2, -4}, 10},
        {[]int{-6, +3, +8, +5, -6}, 5},
        {[]int{+7, +7, -2, -7, -4}, 14},
	}

    for _, table := range tables {
        if got := recurring_1(table.vals); got != table.want {
            t.Errorf("result was incorrect, got: %d, want: %d (%v)", got, table.want, table)
        }
    }
}
