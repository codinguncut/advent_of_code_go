package day4

import (
	"fmt"
	"strconv"
    "aoc"
)

// NumberToDigits converts an integer to a slice of its digits
func NumberToDigits(number int) []int {
	digits := []int{}
	str := strconv.Itoa(number)
	for _, v := range str {
		digit, err := strconv.Atoi(string(v))
        aoc.Check(err)
		digits = append(digits, digit)
	}
	return digits
}


// CalcPreds calculates predicates of a given number
//  that are then used to check whether the number is valid
func CalcPreds(number int) (hasDouble, hasRepeating, isDecreasing bool) {
	digits := NumberToDigits(number)
	last := digits[0]
	digitCounts := map[int]int{last: 1}

	isDecreasing = false
	for _, v := range digits[1:] {
		if v < last {
			isDecreasing = true
		}
		digitCounts[v]++
		last = v
	}

	hasDouble = false
	hasRepeating = false
	for _, v := range digitCounts {
		if v == 2 {
			hasDouble = true
		}
		if v >= 2 {
			hasRepeating = true
		}
	}
	return
}

// CheckValid1 checks if "number" is valid for part 1
func CheckValid1(number int) bool {
    _, hasRepeating, isDecreasing := CalcPreds(number)
    return hasRepeating && !isDecreasing
}

// CheckValid2 checks if "number" is valid for part 2
func CheckValid2(number int) bool {
    hasDouble, _, isDecreasing := CalcPreds(number)
    return hasDouble && !isDecreasing
}

// Main runs the program for day 4
func Main() {
	numFrom := 359282
	numTo := 820401

	count1 := 0
	count2 := 0
	for i := numFrom; i < numTo; i++ {
		if CheckValid1(i) {
			count1++
		}
		if CheckValid2(i) {
			count2++
		}
	}
	fmt.Println("day4.1", count1)
	fmt.Println("day4.2", count2)
}
