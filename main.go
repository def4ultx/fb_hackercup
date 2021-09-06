package main

import "fmt"

func main() {
	answers := make(map[Answer]struct{})

	a1 := Answer{Index: 1, Direction: 0}
	a2 := Answer{Index: 1, Direction: 1}
	a3 := Answer{Index: 5, Direction: 1}
	a4 := Answer{Index: 6, Direction: 0}
	a5 := Answer{Index: 6, Direction: 1}

	answers[a1] = struct{}{}
	answers[a2] = struct{}{}
	answers[a3] = struct{}{}
	answers[a4] = struct{}{}
	answers[a5] = struct{}{}

	fmt.Println(answers)

	fmt.Println("----")

	for key := range answers {
		key.Missing = 0
	}

	fmt.Println(answers)
}

type Answer struct {
	Index     int
	Direction int // 0 for X, 1 for Y
	Missing   int // number of missing X

	// For edge case where missing = 1
	FillX int
	FillY int
}
