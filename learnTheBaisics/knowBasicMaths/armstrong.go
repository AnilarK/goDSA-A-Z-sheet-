package main

import (
	"bufio"
	"fmt"
	"os"
)

func findLen(n int) int {
	var ans int = 0
	for n != 0 {
		ans++
		n = n / 10
	}
	return ans
}

func pow(x int, p int) int {
	var ans int = 1
	for i := 0; i < p; i++ {
		ans = ans * x
	}
	return ans
}

func isArmstrong(n int) bool {
	var orgNum = n
	var len int = findLen(n)
	var armNum int

	for n != 0 {
		orgNum = orgNum + pow(n%10, len)
		n = n / 10
	}

	return orgNum == armNum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	var isArm bool = isArmstrong(n)

	if isArm {
		fmt.Println("This is a Armstrong number")
	} else {
		fmt.Println("This is NOT a Armstrong number")
	}

}
