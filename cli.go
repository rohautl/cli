package main

import (
    "fmt"
    "math/rand"
	"time"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

	min := 1
	max := 10
	numbers := []int{}

	for t := min; t < max + 1; t++{
		numbers = append(numbers, t)
	}
		
	for iterations := len(numbers); iterations > 0; iterations-- {
		i := r1.Intn(len(numbers))
		fmt.Println(numbers[i])

		numbers[i] = numbers[len(numbers)-1]
		numbers = numbers[:len(numbers)-1]
	}
}
