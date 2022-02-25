package main

import (
	"bufio"
	"fmt"
	"os"
)

const apertaenter = "  and press enter when ready."

func main() {

	var firstnumber = 2
	var secondnumber = 5
	var subtraction = 7
	// var answer int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the number game")
	fmt.Println("---------------------------")
	fmt.Println("")

	fmt.Println("think of a number between 1 and 10", apertaenter)
	reader.ReadString('\n')

	fmt.Println("Multply your number by", firstnumber, apertaenter)
	reader.ReadString('\n')

	fmt.Println("Now Multiply the", secondnumber, apertaenter)
	reader.ReadString('\n')

	fmt.Println("devide the resultby the number you originally though of", apertaenter)
	reader.ReadString('\n')

	fmt.Println("Now Subtract", subtraction, apertaenter)
	reader.ReadString('\n')

}
