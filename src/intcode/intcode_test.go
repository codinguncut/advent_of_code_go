package intcode

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
        {[]int{1,9,10,3,2,3,11,0,99,30,40,50}, []int{3500,9,10,70, 2,3,11,0, 99, 30,40,50}},
	}

    for _, table := range tables {
        got := Exec(table.program, nil).memoryArray();
        if !reflect.DeepEqual(got, table.endState) {
            t.Errorf("result was incorrect, got: %v, want: %v",
                got, table.endState)
        }
    }
}

func TestParseOpcode(t *testing.T) {
    raw := Opcode(1002)
    opWant := Opcode(2)
    paramsWant := []ParamMode{0,1,0}
    opGot, paramsGot:= ParseOpcode(raw)
    if opGot != opWant {
        t.Errorf("result was incorrect, got: %d, want: %d", opGot, opWant)
    }
    if !reflect.DeepEqual(paramsWant, paramsGot) {
        t.Errorf("result was incorrect, got: %v, want: %v", paramsGot, paramsWant)
    }
}

func TestParamModes(t *testing.T) {
    tables := []struct {
		program []int
		endState []int
	}{
        {[]int{1002,4,3,4,33}, []int{1002,4,3,4,99}}, // 33*3=99
        {[]int{1101,100,-1,4,0}, []int{1101,100,-1,4,99}}, // 100-1=99
	}
    for _, table := range tables {
        got := Exec(table.program, nil).memoryArray()
        if !reflect.DeepEqual(got, table.endState) {
            t.Errorf("result was incorrect, got: %v, want: %v",
                got, table.endState)
        }
    }
}

func TestDay5(t *testing.T) {
    tables := []struct {
		program []int
		inputs []int
        outputs []int
	}{
        {[]int{3,0,4,0,99}, []int{42}, []int{42}},
        {[]int{3,9,8,9,10,9,4,9,99,-1,8}, []int{7}, []int{0}},
        {[]int{3,9,7,9,10,9,4,9,99,-1,8}, []int{7}, []int{1}},
        {[]int{3,3,1108,-1,8,3,4,3,99}, []int{8}, []int{1}},
        {[]int{3,3,1107,-1,8,3,4,3,99}, []int{8}, []int{0}},
        {[]int{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9}, []int{8}, []int{1}},
        {[]int{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9}, []int{0}, []int{0}},
        {[]int{3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31, 1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104, 999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99}, []int{7}, []int{999}},
    }

    for _, table := range tables {
        got := Exec(table.program, table.inputs).OutputVals
        if !reflect.DeepEqual(got, table.outputs) {
            t.Errorf("result was incorrect, got: %v, want: %v",
                got, table.outputs)
        }
    }
}
