package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Learn the basics ")

	fmt.Println("Enter n : ")
	reader := bufio.NewReader(os.Stdin)
	//input, _ := reader.ReadString('\n')
	//input = strings.TrimSpace(input)
	//n, _ := strconv.Atoi(input)
	//
	//arr := make([]int, 0)
	//for i := 0; i < n; i++ {
	//	input, _ := reader.ReadString(' ')
	//	input = strings.TrimSpace(input)
	//	x, _ := strconv.Atoi(input)
	//	arr = append(arr, x)
	//}
	//
	//fmt.Println(arr)

	var n int
	fmt.Scan(&n)
	brr := make([]int, 0)

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)

		brr = append(brr, x)
	}

	fmt.Println(brr)
	fmt.Println(len(brr))

	var s string
	fmt.Fscan(reader, &s)

	fmt.Println("the string is : ", s)
	fmt.Println("the length of string is : ", len(s))
	fmt.Println("the 3rd character is : ", string(s[2]))

}
