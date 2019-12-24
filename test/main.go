package main

import (
	"fmt"
)

func main() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[2:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) // length of is 2 and capacity is 6
}
