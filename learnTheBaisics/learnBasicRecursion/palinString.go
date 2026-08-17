package main

import (
	"bufio"
	"fmt"
	"os"
)

func valid(s *byte) bool {
	if *s >= 'A' && *s <= 'Z' {
		*s = *s - 'A' + 'a'
	}
	if *s < 'a' || *s > 'z' {
		return false
	}
	return true
}

func isPalin(i int, j int, s string) bool {
	if i >= j {
		return true
	}

	left := s[i]
	right := s[j]

	if !valid(&left) {
		return isPalin(i+1, j, s)
	}

	if !valid(&right) {
		return isPalin(i, j-1, s)
	}

	if left != right {
		return false
	}
	return isPalin(i+1, j-1, s)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the String : ")

	var s string
	fmt.Fscan(reader, &s)

	var n int = len(s)
	var flag bool = isPalin(0, n-1, s)

	if flag {
		fmt.Println("The string is a Palindrome")
	} else {
		fmt.Println("The string is NOT a Palindrome")
	}

}
