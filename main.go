package main

import (
	"fmt"
)

func main() {
	var result string
	var n int
	fmt.Scan(&n)
	fmt.Println("Hello world", n)

	for i := 0; i < n; i++ {
		value := getRow((i + 1), n)
		result += value + "\n"
	}

	fmt.Print(result)
}

func getRow(index int, n int) string {
	var result string
	difference := n - index

	if difference > 0 {
		for i := 0; i < difference; i++ {
			result += " "
		}
	}

	for i := 0; i < index; i++ {
		result += "#"
	}

	return result
}
