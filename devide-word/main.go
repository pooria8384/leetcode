package main

import (
	"fmt"
	"strings"
)

func Gcd(a, b string) int {
	wrd1 := len(a)
	wrd2 := len(b)
	if a+b == b+a {
		for wrd2 != 0 {
			wrd1, wrd2 = wrd2, wrd1%wrd2
		}

	}
	return wrd1
}

func Repeated(str, sub string) bool {
	isRepeated := strings.Repeat(sub, len(str)/len(sub))
	return isRepeated == str
}

func DevideWord(a, b string) string {
	if a+b != b+a {
		fmt.Println("")
		return ""
	}
	gcdLength := Gcd(a, b)
	gcdWord := a[:gcdLength]

	if Repeated(a, gcdWord) && Repeated(b, gcdWord) {
		return gcdWord
	} else {
		fmt.Println("False")
		return ""
	}
}

func main() {
	str1 := "ababab"
	str2 := "ab"
	fmt.Println(DevideWord(str1, str2))

	str1 = "abcdabcd"
	str2 = "abcd"
	fmt.Println(DevideWord(str1, str2))

	str1 = "abcdabcdef"
	str2 = "abcd"
	fmt.Println(DevideWord(str1, str2))

}
