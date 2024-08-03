package main

import "fmt"

func main() {
	name := Reverse("pooria")
	fmt.Println(name)
}
func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
