package main

import (
	"fmt"
)

func main() {
	// pooria()
	hassan()
}

func pooria() {
	word1 := "abcd"
	word2 := "efgh"
	fmt.Println(len(word1))
	empty := " "
	wordLen1 := len(word1) - len(word2)
	wordLen2 := len(word2) - len(word1)
	if len(word1) > len(word2) {
		for i := 0; i < wordLen1; i++ {
			word2 += empty
		}
	} else if len(word2) > len(word1) {
		for i := 0; i < wordLen2; i++ {
			word1 += empty
		}
	}

	for i := range word1 {
		fmt.Printf(string(word1[i]))
		fmt.Printf(string(word2[i]))
	}
}

func hassan() {
	word1 := "abc"
	word2 := "def"
	runeArray1 := []rune(word1)
	runeArray2 := []rune(word2)

	maxLen := max(len(runeArray1), len(runeArray2))

	for i := 0; i < maxLen; i++ {
		w1, w2 := "", ""
		if len(runeArray1) > i {
			w1 = string(runeArray1[i])
		}
		if len(runeArray2) > i {
			w2 = string(runeArray2[i])
		}
		fmt.Printf("%v%v", w1, w2)
	}
}
