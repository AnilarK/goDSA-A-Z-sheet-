package main

import (
	"bufio"
	"fmt"
	"os"
)

func printN(s string, n int) {
	if n <= 0 {
		return
	}
	fmt.Println(s)
	printN(s, n-1)
}

func print1toN(n int) {
	if n <= 0 {
		return
	}
	print1toN(n - 1)
	fmt.Print(n, " ")
}

func printNto1(n int) {
	if n <= 0 {
		return
	}
	fmt.Print(n, " ")
	printNto1(n - 1)
}

func sumOfN(n int) int {
	if n <= 0 {
		return 0
	}
	return n + sumOfN(n-1)
}

func fact(n int) int {
	if n <= 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Learn Recurssion : ")

	fmt.Println("Enter a Number : ")
	var n int
	fmt.Fscan(reader, &n)

	fmt.Println("Enter something : ")
	var s string
	fmt.Fscan(reader, &s)

	printN(s, n)

	print1toN(n)
	fmt.Println()

	printNto1(n)
	fmt.Println()

	var sum int = sumOfN(n)
	fmt.Println("sum of first N number is : ", sum)

	var fac int = fact(n)
	fmt.Println("fact of N number is : ", fac)

}
