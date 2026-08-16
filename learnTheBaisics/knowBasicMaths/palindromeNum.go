package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPalindrome(n int) bool {
	var orgNum int = n
	var revNum int = 0

	for n != 0 {
		x := n % 10
		n = n / 10

		revNum = revNum*10 + x
	}

	return orgNum == revNum
}

func main() {

	fmt.Println("Enter the Number :")
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	var isPalin bool = isPalindrome(n)

	if isPalin {
		fmt.Println("It is a Palindrome")
	} else {
		fmt.Println("It is NOT a Palindrome")
	}

}
