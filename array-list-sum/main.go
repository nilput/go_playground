package main

import (
	"fmt"
	"strings"
)

func printArray(arr []int) {
	if arr == nil {
		fmt.Println("NIL!")
		return
	}
	var sb strings.Builder
	arrayLen := len(arr)
	sb.WriteString("[ ")
	for i, v := range arr {
		sb.WriteString(fmt.Sprintf("%d", v))
		if i+1 < arrayLen {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(" ]")
	fmt.Println(sb.String())
}

func arraySum(arr []int) int {
	if arr == nil {
		return 0
	}
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}

func main() {
	fmt.Println("Array sum program. v1.0")
	fmt.Println("Enter numbers to calculate the sum.")
	var num int
	var arr []int
	for {
		fmt.Print("#: ")
		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			fmt.Errorf("Invalid input")
			break
		}
		arr = append(arr, num)
	}
	printArray(arr)
	fmt.Printf("Sum: %d\n", arraySum(arr))

}
