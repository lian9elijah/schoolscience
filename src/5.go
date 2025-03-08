package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(10) + 1
	fmt.Println("The random number is:", randomNumber)
}
