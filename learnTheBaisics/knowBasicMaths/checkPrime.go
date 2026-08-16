package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPrimeNum(n int) bool {

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Number : ")

	var n int
	fmt.Fscan(reader, &n)

	var isPrime = isPrimeNum(n)

	if isPrime {
		fmt.Println("This is a PRIME NUMBER")
	} else {
		fmt.Println("This is NOT a PRIME NUMBER")
	}

}
