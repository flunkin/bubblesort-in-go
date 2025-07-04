package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bubbleSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	fmt.Println("enter numbers seperated with spaces:")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("issue reading input")
		return
	}

	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	var numbers []int
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("'%s' isnt a real number.\n", part)
			continue
		}
		numbers = append(numbers, num)
	}

	if len(numbers) == 0 {
		fmt.Println("no real numbers entered")
		return
	}

	fmt.Println(numbers)

	bubbleSort(numbers)

	fmt.Println("here is sorted numbers", numbers)
}
