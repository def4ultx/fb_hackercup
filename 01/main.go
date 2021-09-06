package main

import "fmt"

func main() {

	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var s string
		fmt.Scanf("%s", &s)

		count := MaxChange(s)
		fmt.Printf("Case #%d: %d\n", i+1, count)
	}
}

func MaxChange(s string) int {

	var (
		consonants = make(map[rune]int)
		vowels     = make(map[rune]int)

		consonant int
		vowel     int
	)

	for _, v := range s {
		if IsVowel(v) {
			vowels[v]++
			vowel++
		} else {
			consonants[v]++
			consonant++
		}
	}

	maxConsonant := MaxChar(consonants)
	maxVowel := MaxChar(vowels)

	a := (consonant-maxConsonant)*2 + vowel
	b := consonant + (vowel-maxVowel)*2
	result := Min(a, b)

	return result

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IsVowel(r rune) bool {
	return r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func MaxChar(m map[rune]int) int {
	var max = 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}
