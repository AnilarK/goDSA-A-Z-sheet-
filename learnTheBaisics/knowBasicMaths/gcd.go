package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a int, b int) int {
	if a < b {
		return gcd(b, a)
	}

	if b == 0 {
		return a
	}
	if a%b == 0 {
		return b
	}

	return gcd(a%b, b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the First Number : ")

	var firstNum int
	var secondNum int

	fmt.Fscan(reader, &firstNum)
	fmt.Fscan(reader, &secondNum)

	var gcd int = gcd(firstNum, secondNum)
	fmt.Println("the gcd of both number is : ", gcd)
}
