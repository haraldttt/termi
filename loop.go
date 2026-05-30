package main

// cool Go facts - you get compilation errors when having unused imports and variables

import (
	"fmt"
)

func main() {
	var answer string
	correct_answer := "2"
	fmt.Print("What's 1+1?\n")
	fmt.Scanln(&answer)
	if answer == correct_answer {
		fmt.Println("correct answer :)")
	} else {
		fmt.Println("wrong answer :(")
	}

}
