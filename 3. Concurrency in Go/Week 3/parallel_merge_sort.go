// this is not my implementation. this is a submission I reviewed and I'm saving it here for future reference.

package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"
import "log"

func merge(left []int, right []int) []int {
	sorted := make([]int, 0)
	workLeft := make([]int, len(left))
	copy(workLeft, left)
	workRight := make([]int, len(right))
	copy(workRight, right)
	for (len(workLeft) > 0) || (len(workRight) > 0) {
		if len(workLeft) > 0 && len(workRight) > 0 {
			if workLeft[0] <= workRight[0] {
				sorted = append(sorted, workLeft[0])
				workLeft = workLeft[1:]
			} else {
				sorted = append(sorted, workRight[0])
				workRight = workRight[1:]
			}
		} else if len(workLeft) > 0 {
			sorted = append(sorted, workLeft[0])
			workLeft = workLeft[1:]
		} else {
			sorted = append(sorted, workRight[0])
			workRight = workRight[1:]
		}
	}
	return sorted
}

func merge_sort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	middle_index := int(len(array) / 2)
	left := array[0:middle_index]
	right := array[middle_index:]
	left = merge_sort(left)
	right = merge_sort(right)
	return merge(left, right)
}

//The goroutine entry point
func merge_sort_multiple(c chan []int, chunk []int) {
	result := merge_sort(chunk)
	c <- result
}

func parallel_merge_sort(arr []int, workers int) []int {
	step := int(len(arr) / workers)
	chunk := []int{}
	chans := []chan []int{}
	//results := make([][]int, 0)
	for i := 0; i < workers; i++ {
		c := make(chan []int, workers)
		chans = append(chans, c)
		if i < workers-1 {
			chunk = arr[i*step : (i+1)*step]
		} else {
			chunk = arr[i*step:]
		}
		fmt.Println("Chunk", i, "-> ", chunk, " size -> ", len(chunk))

		//start <<workers>> number of go routines for each chunk.
		go merge_sort_multiple(c, chunk)
		//**************
	}
	sorted := make([]int, 0)
	for _, v := range chans {
		sorted = merge(sorted, <-v)
	}
	return sorted

}

func main() {
	fmt.Println("Please enter a list of space separeted integers or CTRL-C to exit")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	strs := strings.Split(line[0:len(line)-1], " ")
	nums := make([]int, len(strs))
	for i, str := range strs {
		if nums[i], err = strconv.Atoi(strings.TrimSpace(str)); err != nil {
			log.Fatal(err)
		}
	}

	result := parallel_merge_sort(nums, 4)
	fmt.Println("Sorted list -> ", result)
}
