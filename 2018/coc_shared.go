package main

// error handling for lazy people
func check(e error) {
    if e != nil {
        panic(e)
    }
}
