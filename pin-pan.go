package main

import (
    "fmt"
    "strconv"
)

func main() {
    for i := 1; i <= 100; i++ {
        switch {
        case i%15 == 0:
            fmt.Println("PinPan")
        case i%3 == 0:
            fmt.Println("Pin")
        case i%5 == 0:
            fmt.Println("Pan")
        default:
            fmt.Println(strconv.Itoa(i))
        }
    }
}
