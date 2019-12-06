package day6

import (
    "testing"
)

func TestCalcPathLens(t *testing.T) {
    lines := []string {
        "COM)B",
        "B)C",
        "C)D",
        "D)E",
        "E)F",
        "B)G",
        "G)H",
        "D)I",
        "E)J",
        "J)K",
        "K)L",
    }
    want := 42

    orbs := makeOrbits()
    orbs.read(lines)
    got := orbs.calcPathLens()
    if got != want {
        t.Errorf("result was incorrect, got: %v, want: %v", got, want)
    }
}

func TestPart2(t *testing.T) {
    lines := []string {
        "COM)B",
        "B)C",
        "C)D",
        "D)E",
        "E)F",
        "B)G",
        "G)H",
        "D)I",
        "E)J",
        "J)K",
        "K)L",
        "K)YOU",
        "I)SAN",
    }
    want := 4

    orbs := makeOrbits()
    orbs.read(lines)
    got := orbs.part2()
    if got != want {
        t.Errorf("result was incorrect, got: %v, want: %v", got, want)
    }
}

