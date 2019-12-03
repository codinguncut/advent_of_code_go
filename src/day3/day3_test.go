package day3

import (
    "testing"
)

func TestDay3(t *testing.T) {
    tables := []struct {
		path1 string
		path2 string
        closest int
        shortest int
	}{
        {"R8,U5,L5,D3", "U7,R6,D4,L4", 6, 30},
        {"R75,D30,R83,U83,L12,D49,R71,U7,L72",
         "U62,R66,U55,R34,D71,R55,D58,R83", 159, 610},
        {"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
         "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135, 410},
	}

    for _, table := range tables {
        p1, p2 := ParsePath(table.path1), ParsePath(table.path2)
        l1, l2 := CalcVisited(p1), CalcVisited(p2)
        got := FindClosest(IntersectVisited(l1, l2))

        if got != table.closest {
            t.Errorf("result was incorrect, got: %d, closest: %d", got,
                     table.closest);
        }

        shortest := FindShortest(l1, l2)
        if shortest != table.shortest{
            t.Errorf("result was incorrect, got: %d, shortest: %d", got,
                     table.shortest);
        }
    }
}
