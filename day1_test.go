package main

import "testing"

func TestFuelRequired(t *testing.T) {
    tables := []struct {
		mass int
		want int
	}{
        {12, 2},
        {14, 2},
        {1969, 654},
        {100756, 33583},
	}

    for _, table := range tables {
        if got := fuel_required(table.mass); got != table.want {
            t.Errorf("result was incorrect, got: %d, want: %d", got, table.want)
        }
    }
}

func TestFuelRequired2(t *testing.T) {
    tables := []struct {
		mass int
		want int
	}{
        {14, 2 + 0},
        {1969, 654 + 216 + 70 + 21 + 5 + 0},
        {100756, 33583 + 11192 + 3728 + 1240 + 411 + 135 + 43 + 12 + 2},
	}

    for _, table := range tables {
        if got := fuel_required_2(table.mass); got != table.want {
            t.Errorf("result was incorrect, got: %d, want: %d", got, table.want)
        }
    }
}

func TestCalcTotalFuel(t *testing.T) {
    // FIXME: reuse table values
    masses := []int{12, 14, 1969, 100756}
    want := 2 + 2 + 654 + 33583
    if got := calc_total_fuel(masses); got != want {
        t.Errorf("result was incorrect, got: %d, want: %d", got, want)
    }
}
