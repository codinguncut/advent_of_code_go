package intcode

import (
    "testing"
    "reflect"
)

func TestExec(t *testing.T) {
    tables := []struct {
		program []int64
		endState []int64
	}{
        {[]int64{1,0,0,0,99}, []int64{2,0,0,0,99}}, // 1+1=2
        {[]int64{2,3,0,3,99}, []int64{2,3,0,6,99}}, // 3*2=6
        {[]int64{2,4,4,5,99,0}, []int64{2,4,4,5,99,9801}}, // 99*99=9801
        {[]int64{1,1,1,4,99,5,6,0,99}, []int64{30,1,1,4,2,5,6,0,99}},
        {[]int64{1,9,10,3,2,3,11,0,99,30,40,50}, []int64{3500,9,10,70, 2,3,11,0, 99, 30,40,50}},
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
    raw := Opcode(21002)
    opWant := Opcode(2)
    paramsWant := []ParamMode{0,1,2}
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
		program []int64
		endState []int64
	}{
        {[]int64{1002,4,3,4,33}, []int64{1002,4,3,4,99}}, // 33*3=99
        {[]int64{1101,100,-1,4,0}, []int64{1101,100,-1,4,99}}, // 100-1=99
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
		program []int64
		inputs []int64
        outputs []int64
	}{
        {[]int64{3,0,4,0,99}, []int64{42}, []int64{42}},
        {[]int64{3,9,8,9,10,9,4,9,99,-1,8}, []int64{7}, []int64{0}},
        {[]int64{3,9,7,9,10,9,4,9,99,-1,8}, []int64{7}, []int64{1}},
        {[]int64{3,3,1108,-1,8,3,4,3,99}, []int64{8}, []int64{1}},
        {[]int64{3,3,1107,-1,8,3,4,3,99}, []int64{8}, []int64{0}},
        {[]int64{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9}, []int64{8}, []int64{1}},
        {[]int64{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9}, []int64{0}, []int64{0}},
        {[]int64{3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31, 1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104, 999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99}, []int64{7}, []int64{999}},
    }

    for _, table := range tables {
        got := Exec(table.program, table.inputs).OutputVals
        if !reflect.DeepEqual(got, table.outputs) {
            t.Errorf("result was incorrect, got: %v, want: %v",
                got, table.outputs)
        }
    }
}

func TestDay9(t *testing.T) {
    tables := []struct {
		program []int64
		inputs []int64
        outputs []int64
	}{
        {
            []int64{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99},
            []int64{},
            []int64{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99},
        },
        {
            []int64{104,1125899906842624,99},
            []int64{},
            []int64{1125899906842624},
        },
    }

    for _, table := range tables {
        got := Exec(table.program, table.inputs).OutputVals
        if !reflect.DeepEqual(got, table.outputs) {
            t.Errorf("result was incorrect, got: %v, want: %v",
                got, table.outputs)
        }
    }
}

func Test203(t *testing.T) {
    program := []int64{9, 1, 203, 0, 99}
    want := []int64{9, 4, 203, 0, 99}
    var wantBase int64 = 1
    inputs := []int64{4}
    got := Exec(program, inputs)
    if int64(got.RelBase) != wantBase {
        t.Errorf("result was incorrect, got: %v, want: %v",
            got.RelBase, wantBase)
    }
    if !reflect.DeepEqual(got.memoryArray(), want) {
        t.Errorf("result was incorrect, got: %v, want: %v",
            got.memoryArray(), want)
    }
}
