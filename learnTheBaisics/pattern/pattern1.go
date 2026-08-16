package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter n")
	var n int
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(j+1, " ")
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(i+1, " ")
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

	for i := n; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

	for i := n; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print(j+1, " ")
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			fmt.Print(" ")
		}
		for j := n - i; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Print("*")
		for j := n - i; j < n; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
	fmt.Println("\n\n\n")

	for i := n - 1; i >= 0; i-- {
		for j := 1; j < n-i; j++ {
			fmt.Print(" ")
		}
		for j := n - i; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Print("*")
		for j := n - i; j < n; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
	fmt.Println("\n\n\n")

	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			fmt.Print(" ")
		}
		for j := n - i; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Print("*")
		for j := n - i; j < n; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
	for i := n - 1; i >= 0; i-- {
		for j := 1; j < n-i; j++ {
			fmt.Print(" ")
		}
		for j := n - i; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Print("*")
		for j := n - i; j < n; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
	fmt.Println("\n\n\n")

	var num int = 1
	var dir bool = true

	for num > 0 {
		for i := 0; i < num; i++ {
			fmt.Print("*")
		}

		if dir {
			num++
		} else {
			num--
		}

		if num == n {
			dir = false
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

	for i := 0; i < n; i++ {
		num = (i + 1) % 2
		for j := 0; j <= i; j++ {
			fmt.Print(num, " ")
			num = 1 - num
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(j + 1)
		}
		for j := i + 1; j < n; j++ {
			fmt.Print(" ")
		}
		for j := i + 1; j < n; j++ {
			fmt.Print(" ")
		}
		for j := i; j >= 0; j-- {
			fmt.Print(j + 1)
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

	num = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(num, " ")
			num++
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(string('A'+j), " ")
			num++
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			fmt.Print(string('A'+j), " ")
			num++
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(string('A'+i), " ")
			num++
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

	for i := 0; i < n; i++ {
		num = 0
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := n - i; j <= n; j++ {
			fmt.Print(string('A' + num))
			num++
		}
		for num > 1 {
			num--
			fmt.Print(string('A' + num - 1))
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

	for i := n; i >= 0; i-- {
		for j := i; j < n; j++ {
			fmt.Print(string('A' + j))
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		for j := i + 1; j < n; j++ {
			fmt.Print(" ")
		}
		for j := i + 1; j < n; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		for j := i + 1; j < n; j++ {
			fmt.Print(" ")
		}
		for j := i + 1; j < n; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("\n\n\n")

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		for j := i + 1; j < n; j++ {
			fmt.Print(" ")
		}
		for j := i + 1; j < n; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		for j := i + 1; j < n; j++ {
			fmt.Print(" ")
		}
		for j := i + 1; j < n; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("\n\n\n")

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == n-1 || j == n-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

	num = n
	for i := 0; i < n+n-1; i++ {
		for j := 0; j < n+n-1; j++ {
			x := min(min(i, j), min(n+n-i-2, n+n-j-2))
			fmt.Print(num-x, " ")
		}
		fmt.Println()
	}
	fmt.Println("\n\n\n")

}
