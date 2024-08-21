package main

import "fmt"

func print_r(messag any) {
    fmt.Print("\033[31m")
    fmt.Print(messag)
    fmt.Println("\033[0m")
}
