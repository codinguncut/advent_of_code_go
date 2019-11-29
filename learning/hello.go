/*
me trying to learn go in a couple of hours.
most of the examples from here: https://gobyexample.com/
*/

package main

import (
    "errors"
    "fmt"
    "sort"
)

func my_println() {
    fmt.Println("Hello World!")
}

func my_variables() {
    var age int = 70
    age = 71
    fmt.Printf("Quantity is %d\n", age)
}

func multi_assign() {
    name, age := "Lemmy", 70
    fmt.Printf("%s's age is %d\n", name, age)
}

func returnMulti() (n int, b bool) {
    n = 42
    b = true
    return
}

func abc() {
    x, _ := returnMulti()
    fmt.Printf("x is %d\n", x)
}

func times(mult int) (func(a int) int) {
    inner := func(a int) int {
        return a * mult
    }
    return inner
}

func test_times() {
    double := times(2)
    fmt.Printf("5 doubled is %d\n", double(5))
}

func constants() {
    const n = 50000
    // n += 1
    fmt.Printf("const is %d\n", n)
}

func loops() {
    for i := 0; i < 3; i++ {
        fmt.Println("for");
    }

    n := 0
    for {
        fmt.Println("forever");
        n += 1
        if n >= 2 {
            break
        }
    }
}

func type_switch() {
    whatAmI := func(i interface{}) {
        switch i.(type) {
        case bool:
            fmt.Println("bool")
        case int:
            fmt.Println("int")
        default:
            fmt.Println("something else")
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}

func arrays() {
    a := [5]int{1,2,3,4,5}
    fmt.Println("array", a)

    b := make([]string, 2)
    b[0] = "hello"
    b[1] = "world"
    b = append(b, "!")
    fmt.Println("strings", b)
    b = append(b, "rubbish")
    fmt.Println("slice", b[:3])
}

func maps() {
    m := make(map[string]int)
    m["k1"] = 7
    m["k2"] = 9
    fmt.Println("map", m)
}

func mult_returns_again () {
    worker := func() (int, int) {
        return 1, 2
    }
    a, b := worker()
    fmt.Println("a", a, "b", b)
}

func variadic_1() {
    nums := []int{1,2,3,4,5}
    worker := func(nums ...int) int {
        total := 0
        for _, num := range nums {
            total += num
        }
        return total
    }

    fmt.Println("variadic sum", nums, worker(nums...))
}

func pointer() {
    worker := func(iptr *int) {
        *iptr = 0
    }
    i := 20
    worker(&i)
    fmt.Println("pointer", i)
}

type person struct {
    name string
    age int
}

func (p person) get_name() string {
    return p.name
}

func my_struct() {
    worker := func(name string) *person {
        p := person{name: name}
        p.age = 42
        return &p
    }

    x := *worker("John")
    fmt.Println("struct", x)
}

func my_method() {
    p := &person{name: "Jonas", age: 42}
    fmt.Println("method", p.get_name())
}

func my_errors() {
    f1 := func(arg int)(int, error) {
        if arg == 42 {
            return -1, errors.New("can't work with 42")
        }
        return arg + 3, nil
    }
    _, err := f1(42)
    if err != nil {
        // panic(err)
    }
}

func my_sort() {
    a := []int{7, 2, 4}
    sort.Ints(a) // in-place
    fmt.Println("sorted", a)
}

/* Notes
- switch without expression
- range is enumerate?
- implicit interfaces?
- go routines
- channels are pipes connecting go routines
- c-style casting
- type assertion inside if `if str, ok := val.(string); ok {...}`
    - no subclassing, only interfaces and struct embedding
*/

func main() {
    my_println()
    my_variables()
    multi_assign()
    abc()
    test_times()
    constants()
    loops()
    type_switch()
    arrays()
    maps()
    mult_returns_again()
    variadic_1()
    pointer()
    my_struct()
    my_errors()
    my_sort()
}

