package main

import (
    "fmt"
    "day1"
    "day2"
    "day3"
    "day4"
    "day5"
    "day6"
    "day7"
    "day8"
    "day9"
    "day10"
)

func main() {
    funcs := [](func()){
        day10.Main,
        day9.Main,
        day8.Main,
        day7.Main,
        day6.Main,
        day5.Main,
        day4.Main,
        day3.Main,
        day2.Main,
        day1.Main,
    }
    for _, f := range funcs {
        f()
        fmt.Println()
    }
}
