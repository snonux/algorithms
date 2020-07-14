package main

import (
	"algorithms/ds"
	"algorithms/sort"
	"fmt"
)

func main() {
	fmt.Println("Hello to Algorithms Playground!")
	a := ds.RandomIntegers(100, 1000)
	fmt.Printf("Random: %v\n", a)
	fmt.Printf("Sorted: %v\n", sort.Insertion(a))
}
