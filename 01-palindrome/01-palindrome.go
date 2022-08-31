package main

import "fmt"

func main() {
	// True or false
	status := false
	fmt.Print("input : ")
	var word string
	fmt.Scanf("%s", &word)

	reverse := func(word string) (result string) {
		for _, v := range word {
			result = string(v) + result
		}
		return
	}

	if word == reverse(word) {
		status = true
	}

	fmt.Println(status)
}
