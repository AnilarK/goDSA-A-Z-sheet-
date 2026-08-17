package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseArrayByRecurssion(i int, j int, a []int) {
	if i >= j {
		return
	}

	var temp int = a[i]
	a[i] = a[j]
	a[j] = temp

	reverseArrayByRecurssion(i+1, j-1, a)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the length of array : ")
	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, 0)
	fmt.Println("Enter elements of array : ")
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(reader, &x)
		a = append(a, x)
	}

	fmt.Println("Array before reversal : ", a)
	reverseArrayByRecurssion(0, n-1, a)
	fmt.Println("Array after reversal : ", a)
}
