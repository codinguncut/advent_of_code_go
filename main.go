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
)

func main() {
    funcs := [](func()){
        day1.Main,
        day2.Main,
        day3.Main,
        day4.Main,
        day5.Main,
        day6.Main,
        day7.Main,
    }
    for _, f := range funcs {
        f()
        fmt.Println()
    }
}
