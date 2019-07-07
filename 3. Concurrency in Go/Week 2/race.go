// Write two goroutines which have a race condition when executed concurrently.
// Explain what the race condition is and how it can occur.

package main

import (
	"fmt"
)

func main() {

	// We are changing the value of i in 2 separate go routines, to 1 and 2 (default is 0).
	// Now depending on which of these go routines completes first, our value printed will be
	// either 0 (the default integer value) or 1 or 2.
	// This is the race condition, the value of i changes depending on which go routine completes first.
	// I print the value of 2 three times, just to see how it changes with time.

	i := 0
	go func() {
		i = 1
		fmt.Println("i = ", i, " completed")
	}()
	go func() {
		i = 2
		fmt.Println("i = ", i, " completed")
	}()

	for j := 0; j < 3; j++ {
		fmt.Println(i)
	}
}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	// use WaitGroup to ensure all the goroutines finish
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	// There is a race condition here because the variable
// 	// 'i' is referenced from within the anonymous function.
// 	// This results in the read from the goroutine racing against the
// 	// increment of 'i' as part of the for loop.
// 	//
// 	for i := 0; i < 2; i++ {
// 		go func() {
// 			fmt.Println(i)
// 			wg.Done()
// 		}()
// 	}

// 	wg.Wait()
// }
