package main

import (
	"bufio"
	"fmt"
	"os"
)

func rev(num int) int {
	var revNum int

	for num != 0 {
		lastDigit := num % 10
		num = num / 10
		revNum = revNum*10 + lastDigit
	}

	return revNum
}

func main() {

	fmt.Println("Enter a number :")

	reader := bufio.NewReader(os.Stdin)

	var num int
	fmt.Fscan(reader, &num)

	fmt.Println("the reversed number is : ", rev(num))

}
