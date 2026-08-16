package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Learn if-else :")

	reader := bufio.NewReader(os.Stdin)

	var age int
	fmt.Fscan(reader, &age)

	if age > 18 {
		fmt.Println("you are a adult")
	} else {
		fmt.Println("you are not a adult")
	}

}
