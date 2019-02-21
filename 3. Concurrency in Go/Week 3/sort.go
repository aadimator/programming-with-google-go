// Write a program to sort an array of integers. The program should partition the array
// into 4 parts, each of which is sorted by a different goroutine. Each partition should
// be of approximately equal size. Then the main goroutine should merge the 4 sorted
// subarrays into one large sorted array.

// The program should prompt the user to input a series of integers. Each goroutine which
// sorts ¼ of the array should print the subarray that it will sort. When sorting is complete,
// the main goroutine should print the entire sorted list.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// Extracts integers from a string
func stringToInts(s string) []int {
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
	fmt.Print("Enter a list of integers separated by space: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nums := stringToInts(scanner.Text())

	const partition = 4
	n := int(math.Max(math.Ceil(float64(len(nums))/float64(partition)), 1))

	var wg sync.WaitGroup

	// partition the given array to approximately equal sizes
	for i := 0; i < partition; i++ {
		from := int(math.Min(float64(i*n), float64(len(nums))))
		to := int(math.Min(float64((i+1)*n), float64(len(nums))))

		wg.Add(1)

		go func(arr []int) {
			fmt.Println("This go routine will sort: ", arr)
			sort.Ints(arr)

			wg.Done()
		}(nums[from:to])
	}

	// wait for all goroutines to be finished
	wg.Wait()

	sort.Ints(nums)
	fmt.Println("\nSorted List: ", nums)
}
