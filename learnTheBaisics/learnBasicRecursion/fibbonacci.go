package main

import (
	"bufio"
	"fmt"
	"os"
)

func fibbo(n int) int {
	if n <= 1 {
		return n
	}
	return fibbo(n-1) + fibbo(n-2)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the number : ")

	var n int
	fmt.Fscan(reader, &n)

	var fib int = fibbo(n)

	fmt.Println("The nth Fibbonacci number is : ", fib)

}
