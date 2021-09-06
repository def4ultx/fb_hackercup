package main_test

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func main() {

	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		runtest(i)
	}

	// runtest(0)
}

func runtest(tc int) {
	var str string
	fmt.Scanf("%s", &str)

	var n int
	fmt.Scanf("%d", &n)

	var (
		replacements = make([]string, n)
		mapping      = make(map[rune][]rune)
		destinations = make(map[rune]struct{})
		costs        = make(map[string]int)
	)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &replacements[i])

		src, _ := utf8.DecodeRune([]byte{replacements[i][0]})
		dst, _ := utf8.DecodeRune([]byte{replacements[i][1]})
		if _, ok := mapping[src]; ok {
			mapping[src] = append(mapping[src], dst)
		} else {
			mapping[src] = []rune{dst}
		}
		destinations[dst] = struct{}{}
		costs[replacements[i]] = 1
	}

	// need to check if all same character
	isSameRune := true
	char, _ := utf8.DecodeRune([]byte{str[0]})
	for _, v := range str {
		if v != char {
			isSameRune = false
			break
		}
	}

	if isSameRune {
		// fmt.Println("edge case, same character")
		fmt.Println("result test case", str, 0)
		return
	}

	// fmt.Println(mapping)

	// fmt.Println("mapping")
	// for k, v := range mapping {
	// 	fmt.Print(string(k) + ":")
	// 	fmt.Println(v)
	// 	for _, s := range v {
	// 		fmt.Print(" " + string(s))
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println(findCost('A', 'C', mapping, costs, false, map[rune]struct{}{}))

	// fmt.Println("costs", costs)
	// fmt.Println("destinations", destinations)
	// fmt.Println(replacements)

	// fmt.Println("result")
	// fmt.Println(findCost('A', 'B', mapping, costs))

	// fmt.Println("----")

	var result []int
	for dst := range destinations {

		// fmt.Println(string(dst))

		var total int
		for _, v := range str {
			// try replace v with dst
			// fmt.Println(string(v), " to char", string(dst))

			if v == dst {
				continue
			}

			cost := findCost(v, dst, mapping, costs, false, map[rune]struct{}{})

			// fmt.Println(cost, total)

			if cost == -1 {
				total = cost
				break
			} else {
				total += cost
			}

			// fmt.Println(cost, total)
			// if cost == math.MaxInt32 {
			// 	break
			// }
		}

		if total != -1 {
			result = append(result, total)
		}

		// fmt.Println("----")
	}

	// fmt.Println("results", result)

	min := math.MaxInt32
	for _, v := range result {
		if v < min {
			min = v
		}
	}

	if min == math.MaxInt32 {
		min = -1
	}

	fmt.Println("result test case", str, min)

}

type Pair struct {
	Character rune
	Cost      int
}

func findCost(src, dst rune, mapping map[rune][]rune, costs map[string]int, canMap bool, cyclic map[rune]struct{}) int {

	key := string(src) + string(dst)
	cost, ok := costs[key]
	if ok {
		return cost
	}

	runes, ok := mapping[src]
	if !ok {

		// fmt.Println("return not found")
		costs[key] = -1
		return -1
	}

	_, ok = cyclic[src]
	if ok {
		// fmt.Println("return not found 3, cyclic")
		costs[key] = -1
		return -1
	}

	newCyclic := make(map[rune]struct{})
	newCyclic[src] = struct{}{}
	for k, v := range cyclic {
		newCyclic[k] = v
	}

	// Need cyclic detection ??
	pairs := make([]Pair, 0)
	for _, r := range runes {
		c := findCost(rune(r), dst, mapping, costs, true, newCyclic)

		if c != -1 {
			pair := Pair{
				Character: rune(r),
				Cost:      c,
			}
			pairs = append(pairs, pair)
		}
	}

	// first find cannot be mapping !!
	if len(pairs) == 0 {
		// fmt.Println("return not found 2")
		return -1 // Not found
	}

	min := pairs[0].Cost
	for _, v := range pairs {
		if v.Cost < min {
			min = v.Cost
		}
	}

	if canMap {

		costs[key] = min + 1
		// fmt.Println("create mapping", key, costs[key])
	}

	return min + 1
}
