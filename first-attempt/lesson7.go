/*
Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence
of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest.

Use your favorite search tool to find a description of how the bubble sort algorithm works.
As part of this program, you should write a function called BubbleSort() which takes a slice
of integers as an argument and returns nothing.

The BubbleSort() function should modify the slice so that the elements are in sorted order.
A recurring operation in the bubble sort algorithm is the Swap operation which swaps the
position of two adjacent elements in the slice.

You should write a Swap() function which performs this operation.
Your Swap() function should take two arguments, a slice of integers
and an index value i which indicates a position in the slice.

The Swap() function should return nothing, but it should swap the contents of the slice in
position i with the contents in position i+1.

Testing:
Test the program by running it twice and
testing it with a different sequence of integers each time. The first test
sequence of integers should be all positive numbers and the second test should
have at least one negative number.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a sequence of up to 10 integers")
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	numbers := parseInput(input)

	if len(numbers) > 10 {
		fmt.Println("Error: Please enter up to 10 integers.")
		return
	}

	fmt.Println("Original sequence:", numbers)
	BubbleSort(numbers)
	fmt.Println("Sorted sequence:", numbers)
}

func parseInput(input string) []int {
	numStrings := strings.Fields(input)
	numbers := make([]int, len(numStrings))

	for i, numStr := range numStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error: Invalid input. Please enter integers only.")
			numbers = nil
			break
		}
		numbers[i] = num
	}

	return numbers
}

func BubbleSort(numbers []int) {
	n := len(numbers)
	swapped := true

	for swapped {
		swapped = false

		for i := 0; i < n-1; i++ {
			if numbers[i] > numbers[i+1] {
				Swap(numbers, i)
				swapped = true
			}
		}

		n--
	}
}

func Swap(numbers []int, i int) {
	temp := numbers[i]
	numbers[i] = numbers[i+1]
	numbers[i+1] = temp
}
