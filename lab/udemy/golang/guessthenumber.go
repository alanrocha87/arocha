package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// var firstnumber = 2
	// var secondnumber = 5
	// var subtraction = 7
	// var answer int
	// var answer int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the number game")
	fmt.Println("---------------------------")
	fmt.Println("")

	fmt.Println("think of a number between 1 and 10 and press ENTER")
	reader.ReadString('\n')

}
