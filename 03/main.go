package main

import (
	"fmt"
	"math"
)

func main() {

	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		runtest(i)
	}
}

type Answer struct {
	Index     int
	Direction int // 0 for X, 1 for Y
	Missing   int // number of missing X

	// For edge case where missing = 1
	FillX int
	FillY int
}

func runtest(tc int) {

	n, original, filled := fillBoard()
	solutions := findAllSolution(n, filled)
	solutionsWithMissing := addMissing(n, original, solutions)

	var min int = math.MaxInt32
	for key := range solutionsWithMissing {
		if key.Missing < min {
			min = key.Missing
		}
	}

	var count int
	switch {
	case min == math.MaxInt32:
	case min != 1:
		count = countMin(min, solutionsWithMissing)
	default:
		count = countMinEdgeCase(n, original, min, solutionsWithMissing)
	}

	if min != math.MaxInt32 {
		fmt.Printf("Case #%d: %d %d\n", tc+1, min, count)
	} else {
		fmt.Printf("Case #%d: Impossible\n", tc+1)
	}
	return
}

func fillBoard() (int, [][]rune, [][]rune) {
	var n int
	fmt.Scanf("%d", &n)

	var (
		original = make([][]rune, n)
		filled   = make([][]rune, n)
	)

	for i := 0; i < n; i++ {
		original[i] = make([]rune, n)
		filled[i] = make([]rune, n)
	}

	for i := 0; i < n; i++ {
		var s string
		fmt.Scanf("%s", &s)

		for j, r := range s {
			original[i][j] = r
			if r == '.' {
				filled[i][j] = 'X'
			} else {
				filled[i][j] = r
			}
		}
	}

	return n, original, filled
}

func findAllSolution(n int, filled [][]rune) map[Answer]struct{} {
	solutions := make(map[Answer]struct{})
	for i := 0; i < n; i++ {

		var (
			foundY = true
			foundX = true
		)

		for j := 0; j < n; j++ {
			if filled[i][j] != 'X' {
				foundX = false
				break
			}
		}

		for j := 0; j < n; j++ {
			if filled[j][i] != 'X' {
				foundY = false
				break
			}
		}

		if foundX {
			s := Answer{
				Index:     i,
				Direction: 0,
			}
			solutions[s] = struct{}{}
		}
		if foundY {
			s := Answer{
				Index:     i,
				Direction: 1,
			}
			solutions[s] = struct{}{}
		}

		// fmt.Println(foundX, foundY)
	}

	return solutions
}

func addMissing(n int, original [][]rune, solutions map[Answer]struct{}) map[Answer]struct{} {
	solutionsWithMissing := make(map[Answer]struct{})
	for s := range solutions {
		var count int
		if s.Direction == 0 {
			for i := 0; i < n; i++ {
				if original[s.Index][i] == '.' {
					count++
				}
			}
		} else {
			for i := 0; i < n; i++ {
				if original[i][s.Index] == '.' {
					count++
				}
			}
		}
		s.Missing = count
		solutionsWithMissing[s] = struct{}{}
	}

	return solutionsWithMissing
}

func countMin(min int, solutions map[Answer]struct{}) int {
	var count int
	for key := range solutions {
		if key.Missing == min {
			count++
		}
	}
	return count
}

func countMinEdgeCase(n int, original [][]rune, min int, solutions map[Answer]struct{}) int {

	type Ans struct {
		X, Y int
	}

	// sol := make(map[Ans]struct{})

	sol := make(map[Answer]struct{})
	for s := range solutions {
		if s.Missing == min {
			if s.Direction == 0 {
				for i := 0; i < n; i++ {
					if original[s.Index][i] == '.' {
						s.FillX = s.Index
						s.FillY = i
						// s.Direction = 0
						// s.Index = 0
						// s.Missing = 0

						// ans := Ans{s.Index, i}
						// sol[ans] = struct{}{}
					}
				}
			} else {
				for i := 0; i < n; i++ {
					if original[i][s.Index] == '.' {
						s.FillX = i
						s.FillY = s.Index
						// s.Direction = 0
						// s.Index = 0
						// s.Missing = 0

						// ans := Ans{i, s.Index}
						// sol[ans] = struct{}{}
					}
				}
			}
			s.Direction = 0

			s.Direction = 0
			s.Index = 0
			s.Missing = 0

			sol[s] = struct{}{}
		}
	}
	fmt.Println(sol)
	return len(sol)
}
