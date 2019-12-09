package day7

import (
    "testing"
    "reflect"
)

func TestPermutations(t *testing.T) {
    tables := []struct {
		inputs []int64
		want []([]int64)
	}{
        {[]int64{1,2}, [][]int64{{1,2}, {2,1}}},
        {[]int64{1,2,3}, [][]int64{{1,2,3}, {1,3,2}, {2,1,3}, {2,3,1}, {3,1,2}, {3,2,1}}},
	}

    for _, table := range tables {
        got := permutations(table.inputs)
        if !reflect.DeepEqual(got, table.want) {
            t.Errorf("result was incorrect, got: %v, want: %v",
                got, table.want)
        }
    }
}

func TestCheckPerm(t *testing.T) {
    tables := []struct {
        program []int64
		perm []int64
		want int64
	}{
        {
            []int64{3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0},
            []int64{4,3,2,1,0},
            43210,
        },
        {
            []int64{3,23,3,24,1002,24,10,24,1002,23,-1,23, 101,5,23,23,1,24,23,23,4,23,99,0,0},
            []int64{0,1,2,3,4},
            54321,
        },
        {
            []int64{3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33, 1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0},
            []int64{1,0,4,3,2},
            65210,
        },
    }
    for _, table := range tables {
        got := checkPerm(table.program, table.perm)
        if got != table.want {
            t.Errorf("result was incorrect, got: %v, want: %v",
                got, table.want)
        }
    }
}

func TestFeedbackLoop(t *testing.T) {
    tables := []struct {
        program []int64
		perm []int64
		want int64
	}{
        {
            []int64{3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26, 27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5},
            []int64{9,8,7,6,5},
            139629729,
        },
        {
            []int64{3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54, -5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4, 53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10},
            []int64{9,7,8,5,6},
            18216,
        },
    }
    for _, table := range tables {
        got := checkPermFeedback(table.program, table.perm)
        if got != table.want {
            t.Errorf("result was incorrect, got: %v, want: %v",
                got, table.want)
        }
    }
}
