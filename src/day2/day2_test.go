package day2

import (
    "testing"
    "reflect"
)

func TestExec(t *testing.T) {
    tables := []struct {
		program []int
		endState []int
	}{
        {[]int{1,0,0,0,99}, []int{2,0,0,0,99}}, // 1+1=2
        {[]int{2,3,0,3,99}, []int{2,3,0,6,99}}, // 3*2=6
        {[]int{2,4,4,5,99,0}, []int{2,4,4,5,99,9801}}, // 99*99=9801
        {[]int{1,1,1,4,99,5,6,0,99}, []int{30,1,1,4,2,5,6,0,99}},
	}

    for _, table := range tables {
        if got := Exec(table.program).pos; !reflect.DeepEqual(got, table.endState) {
            t.Errorf("result was incorrect, got: %d, want: %d", got, table.endState)
        }
    }
}
