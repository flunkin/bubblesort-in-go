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
        didSwap := false
        for j := 0; j < n-1; j++ {
            if numbers[j] > numbers[j+1] {
                temp := numbers[j]
                numbers[j] = numbers[j+1]
                numbers[j+1] = temp
                didSwap = true
            }
        }
        if didSwap == false {
            break
        }
    }
}

func main() {
    fmt.Println("Enter numbers separated by spaces:")

    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input")
        return
    }
    input = strings.TrimSpace(input)

    parts := strings.Split(input, " ")

    var numbers []int
    for i := 0; i < len(parts); i++ {
        num, err := strconv.Atoi(parts[i])
        if err != nil {
            fmt.Println(parts[i], "is not a valid number, skipping...")
            continue
        }
        numbers = append(numbers, num)
    }

    if len(numbers) == 0 {
        fmt.Println("No valid numbers entered, exiting.")
        return
    }

    fmt.Println("Numbers before sorting:", numbers)

    bubbleSort(numbers)

    fmt.Println("Numbers after sorting:", numbers)
}
