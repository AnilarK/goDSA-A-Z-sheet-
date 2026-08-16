package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter a number :")

	reader := bufio.NewReader(os.Stdin)

	var num int
	fmt.Fscan(reader, &num)

	var revNum int

	for num != 0 {
		lastDigit := num % 10
		num = num / 10
		revNum = revNum*10 + lastDigit
	}

	fmt.Println("the reversed number is : ", revNum)

}
