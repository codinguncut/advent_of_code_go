package day4

import (
    "testing"
)

func TestPart1(t *testing.T) {
    tables := []struct {
        number int
        isValid bool
	}{
	    {111111, true},
	    {223450, false},
	    {123789, false},
        {111123, true},
	    {134559, true},
	}

    for _, table := range tables {
        if got := CheckValid1(table.number); got != table.isValid{
            t.Errorf("result was incorrect, got: %v, closest: %v", got,
                     table.isValid);
        }
    }
}

func TestPart2(t *testing.T) {
    tables := []struct {
        number int
        isValid bool
	}{
	    {112233, true},
	    {123444, false},
	    {111122, true},
	}

    for _, table := range tables {
        if got := CheckValid2(table.number); got != table.isValid{
            t.Errorf("result was incorrect, got: %v, closest: %v", got,
                     table.isValid);
        }
    }
}
