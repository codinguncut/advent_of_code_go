package day8

import (
    "fmt"
    "aoc"
    "strconv"
    "strings"
)

// Layer is
type Layer struct {
    // TODO: use 2d slice which knows its width and height?
    data []int
    width int
    height int
}

// color constants
const (
    NumBlack = 0
    NumWhite = 1
    NumTrans = 2
)

func imageToLayers(width, height int, data []int) (layers []Layer) {
    mult := width * height
    for offset := 0; len(data[offset:]) > 0; offset += mult {
        layers = append(layers, Layer{
            data[offset:offset+mult],
            width,
            height,
        })
    }
    return
}

func countDigits(layers []Layer) (res [](map[int]int)) {
    for _, layer := range layers {
        counts := map[int]int{}
        for _, v := range layer.data {
            counts[v]++
        }
        res = append(res, counts)
    }
    return
}

func readLayers(width, height int, str string) (layers []Layer) {
    data := []int{}
    for _, r := range str {
        val, err := strconv.Atoi(string(r))
        aoc.Check(err)
        data = append(data, val)
    }
    return imageToLayers(width, height, data)
}

func printLayer(layer Layer) {
    width := layer.width
    for row := 0; row < len(layer.data) / width; row++ {
        vals := layer.data[row*width:(row+1)*width]
        for _, v := range vals {
            if v == NumBlack {
                fmt.Print(" ")
            } else {
                fmt.Print("#")
            }
        }
        fmt.Println()
    }
}

func stackLayers(layers []Layer) (output Layer) {
    first := layers[0]
    output = Layer{
        append([]int(nil), first.data...),
        first.width,
        first.height,
    }
    for _, layer := range layers[1:] {
        for i := range layer.data {
            if output.data[i] == NumTrans {
                output.data[i] = layer.data[i]
            }
        }
    }
    return
}

func part1(layers []Layer) int {
    counts := countDigits(layers)
    first := counts[0]

    minZero := first[0]
    val := first[1]*first[2]
    for _, cts := range counts[1:] {
        if cts[0] < minZero {
            minZero = cts[0]
            val = cts[1]*cts[2]
        }
    }
    return val
}

func part2(layers []Layer) {
    res := stackLayers(layers)
    printLayer(res)
}


// Main is
func Main() {
    str := strings.TrimSpace(aoc.ReadFile("data/day8_input.txt"))
    layers := readLayers(25, 6, str)

    fmt.Println("day8.1", part1(layers))
    part2(layers)
}
