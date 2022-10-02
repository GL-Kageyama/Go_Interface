package main

import (
    "fmt"
    "./dog"
)

func main() {
    dog := dog.NewDog()
    fmt.Println(dog.GetDog()) // Bow Wow !
}