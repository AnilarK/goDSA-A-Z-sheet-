package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func allDivisor(n int) []int {
	arr := make([]int, 0)

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			arr = append(arr, i)

			if i*i != n {
				arr = append(arr, n/i)
			}
		}
	}

	sort.Ints(arr)
	return arr
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number : ")

	var n int
	fmt.Fscan(reader, &n)

	var arr = allDivisor(n)
	fmt.Println(arr)
}
