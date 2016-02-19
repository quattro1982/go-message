package main

import (
    "fmt"
)

func ball(from string) {
    for i := 0; i < 10; i++{
        fmt.Println(from, ":", i)
    }

}

func main() {
    
    go ball("routine1")
    ball("Main")
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}