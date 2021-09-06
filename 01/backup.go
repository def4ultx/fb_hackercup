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

	// MaxChange("ABC")
	// MaxChange("F")
	// MaxChange("BANANA")
	// MaxChange("FBHC")
	// MaxChange("FOXEN")
	// MaxChange("CONSISTENCY")

	// Case #1: 2
	// Case #2: 0
	// Case #3: 3
	// Case #4: 4
	// Case #5: 5
	// Case #6: 12
}

func MaxChange(s string) int {

	var (
		consonants = make(map[rune]int)
		vowels     = make(map[rune]int)

		consonantCount int
		vowelCount     int
	)

	var result int

	for _, v := range s {
		if isVowel(v) {
			vowels[v]++
			vowelCount++
		} else {
			consonants[v]++
			consonantCount++
		}
	}

	maxConsonant := maxLetter(consonants)
	maxVowel := maxLetter(vowels)

	// fmt.Println(maxConsonant, maxVowel, consonantCount, vowelCount, maxVowel, vowelCount/2)

	// consonant >= vowel && max_consonant <= consonant/2
	// => change to vowel
	// consonant >= vowel && max_consonant > consonant/2
	// => change to consonant

	// consonant < vowel && max_vowel <= vowel/2
	// => change to consonant
	// consonant < vowel && max_vowel > vowel/2
	// => change to vowel

	// FOXEN => OOOEO => OOOBO => OOOOO

	if consonantCount > vowelCount && maxConsonant*2 <= consonantCount {

		result = consonantCount + (vowelCount-maxVowel)*2

	} else if consonantCount > vowelCount && maxConsonant*2 > consonantCount {

		result = (consonantCount-maxConsonant)*2 + vowelCount

	} else if consonantCount < vowelCount && maxVowel*2 <= vowelCount {

		result = (consonantCount-maxConsonant)*2 + vowelCount

	} else if consonantCount < vowelCount && maxVowel*2 > vowelCount {

		result = consonantCount + (vowelCount-maxVowel)*2

	} else if consonantCount == vowelCount && maxVowel <= maxConsonant {

		result = (consonantCount-maxConsonant)*2 + vowelCount

	} else if consonantCount == vowelCount && maxVowel > maxConsonant {

		result = consonantCount + (vowelCount-maxVowel)*2

	} else {

	}

	// fmt.Println(result)

	return result

}

func isVowel(r rune) bool {
	return r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func maxLetter(m map[rune]int) int {
	var max = 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

// ABC > AAA
// consonant = 2
// max_consonant = 1
// vowel = 1
// max_vowel = 1

// ans = 2

// F > F
// consonant = 1
// max_consonant = 1
// vowel = 0
// max_vowel = 0

// ans = 0

// BANANA =>
// consonant = 3
// max_consonant = 1
// vowel = 3
// max_vowel = 3

// ans = 3
// AAAAAA

// FBHC => FFFF
// consonant = 4
// max_consonant = 1
// vowel = 0
// max_vowel = 0

// ans = 4
// AAAAAA

// FOXEN => OOOEO => OOOBO => OOOOO

// consonant = 3
// max_consonant = 1
// vowel = 2
// max_vowel = 1

// ans = 5

// CONSISTENCY

// consonant = 8
// max_consonant = 2
// vowel = 3
// max_vowel = 1

// C O N S I S T E N C Y
// C C N S C S T C N C Y => 3
// C C A A C A A C A C A => 3 + 6
// C C C C C C C C C C C => 3 + 6 + 6
// 15

// ans = 12

// consonant >= vowel && max_consonant <= consonant/2
// => change to vowel

// consonant >= vowel && max_consonant > consonant/2
// => change to consonant

// consonant < vowel && max_vowel <= vowel/2
// => change to consonant
// consonant < vowel && max_vowel > vowel/2
// => change to vowel
