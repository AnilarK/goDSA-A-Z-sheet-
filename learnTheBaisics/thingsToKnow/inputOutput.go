package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Doing input and output")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	num := strings.TrimSpace(input)
	fmt.Println("the input number is : ", num)

}
