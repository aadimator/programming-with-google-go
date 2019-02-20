// Write a Bubble Sort program in Go. The program should prompt the user to
// type in a sequence of up to 10 integers. The program should print the integers
// out on one line, in sorted order, from least to greatest. Use your favorite search
// tool to find a description of how the bubble sort algorithm works.

// As part of this program, you should write a function called BubbleSort() which takes
// a slice of integers as an argument and returns nothing. The BubbleSort() function should
// modify the slice so that the elements are in sorted order.

// A recurring operation in the bubble sort algorithm is the Swap operation which swaps the
// position of two adjacent elements in the slice. You should write a Swap() function which
// performs this operation. Your Swap() function should take two arguments, a slice of integers
// and an index value i which indicates a position in the slice. The Swap() function should return
// nothing, but it should swap the contents of the slice in position i with the contents in position i+1.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Swap function swaps two adjacent elements in a slice
func Swap(s []int, i int) {
	s[i], s[i+1] = s[i+1], s[i]
}

// BubbleSort sorts the slice, from least to greatest.
func BubbleSort(s []int) {
	swapped := false
	for i := 0; i < len(s)-1; i++ {
		swapped = false

		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				Swap(s, j)
				swapped = true
			}
		}

		if !swapped {
			break
		}

	}
}

// Extracts integers from a string
func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func main() {

	fmt.Println("Enter a list of numbers (Enter to stop)")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nums := numbers(scanner.Text())
	BubbleSort(nums)

	fmt.Println("\nSorted List")
	fmt.Println(nums)
}
