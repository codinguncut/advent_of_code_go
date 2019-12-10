package day10

import (
    "testing"
    "reflect"
    "math"
)

var vecs = readLines([]string{
    ".#..#",
    ".....",
    "#####",
    "....#",
    "...##",
})

func TestReadLines(t *testing.T) {
    want := []vector{
        vector{1,0},
        vector{4,0},
        vector{0,2},
        vector{1,2},
        vector{2,2},
        vector{3,2},
        vector{4,2},
        vector{4,3},
        vector{3,4},
        vector{4,4},
    }
    // TODO: not using custom "equal"
    if got := vecs; !reflect.DeepEqual(got, want) {
        t.Errorf("result was incorrect, got: %v, want: %v", got, want)
    }
}

func TestCountVisible(t *testing.T) {
    want := []int{7,7,6,7,7,7,5,7,8,7}
    if got := countVisible(vecs); !reflect.DeepEqual(got, want) {
        t.Errorf("result was incorrect, got: %v, want: %v", got, want)
    }
}

func TestPart1(t *testing.T) {
    wantVal, wantVec := 8, vector{3, 4}
    gotVal, gotIdx := part1(vecs)
    gotVec := vecs[gotIdx]
    if gotVal != wantVal {
        t.Errorf("gotVal, result was incorrect, got: %v, want: %v", gotVal, wantVal)
    }

    if !reflect.DeepEqual(gotVec, wantVec) {
        t.Errorf("gotVec, result was incorrect, got: %v, want: %v",
            gotVec, wantVec)
    }
}

func TestVectorCmp(t *testing.T) {
    v1 := vector{0.3, 0.4}
    v2 := vector{0.300005, 0.3999999}
    if got := v1.equal(v2); !got {
        t.Errorf("got: %v, want: %v", got, true)
    }

    v1 = vector{0.3, 0.4}
    v2 = vector{0.31, 0.4}
    if got := v1.equal(v2); got {
        t.Errorf("got: %v, want: %v", got, false)
    }
}

func TestPolar(t *testing.T) {
    vector{3, 0}.polar()
    vector{0, 4}.polar()
    vector{-3, 0}.polar()

    vec := vector{0, -4}
    want := polar{0, 4}
    if got := vec.polar(); !got.equal(want) {
        t.Errorf("got: %v, want: %v", got, want)
    }

    vec2 := vector{3, -5}
    want2 := polar{math.Pi/2.0 - 1.03037682652431250574, 5.83095189484530}
    if got2 := vec2.polar(); !got2.equal(want2) {
        t.Errorf("got2: %v, want2: %v", got2, want2)
    }
}

func TestReverse(t *testing.T) {
    pol := polar{math.Pi/2.0 - 1.03037682652431250574, 5.83095189484530}
    want := vector{3, -5}
    if got := pol.vector(); !got.equal(want) {
        t.Errorf("got: %v, want: %v", got, want)
    }
}

func TestVectorMag(t *testing.T) {
    got := vector{1, 2}.magnitude()
    want := floatEqual(got, math.Sqrt(5))
    if !want {
        t.Errorf("got: %v, want: %v", got, want)
    }

    got2 := vector{3, 0}.unit()
    want2 := vector{1, 0}
    if !got2.equal(want2) {
        t.Errorf("got: %v, want: %v", got2, want2)
    }
}

func TestDropMany(t *testing.T) {
    xs := []polar{
        polar{1,2},
        polar{2,3},
        polar{3,4},
    }
    y := dropMany(xs, map[int]bool{0: true, 2: true})
    if len(y) != 1 || !y[0].equal(xs[1]) {
        t.Errorf("got: %v, want, %v, %v", 1, len(y), y)
    }
}
