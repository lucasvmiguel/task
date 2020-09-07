package main

import (
	"fmt"

	"github.com/lucasvmiguel/task/internal/calculator"
)

func main() {
	sum := calculator.Sum(2, 2)
	fmt.Println(sum)
	fmt.Println("Hello, world.")
}
