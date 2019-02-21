// Implement the dining philosopher’s problem with the following constraints/modifications.
//
// 1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
// 2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
// 3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
// 4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
// 5. The host allows no more than 2 philosophers to eat concurrently.
// 6. Each philosopher is numbered, 1 through 5.
// 7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>”
//    on a line by itself, where <number> is the number of the philosopher.
// 8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>”
//    on a line by itself, where <number> is the number of the philosopher.

package main

import (
	"fmt"
	"sync"
)

type Chopstick struct {
}

// Each philosopher is numbered, 1 through 5
type Philosopher struct {
	num int
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(Chopstick)
	},
}

var wg sync.WaitGroup

func (p Philosopher) eat(host chan int) {
	defer wg.Done()

	// get permission from the host
	<-host

	fmt.Printf("starting to eat %d\n", p.num)

	// pick up chopsticks in any order
	left := pool.Get()
	right := pool.Get()

	// then return the chopsticks
	pool.Put(left)
	pool.Put(right)

	fmt.Printf("finishing eating %d\n", p.num)

	host <- 1
}

func main() {

	// host allows no more than 2 philosophers to eat concurrently
	host := make(chan int, 2)

	// initialize chopsticks pool
	for i := 0; i < 5; i++ {
		pool.Put(new(Chopstick))
	}

	// initialize philosophers
	philosophers := make([]*Philosopher, 5)

	for i := 0; i < 5; i++ {
		// each philosopher is numbered, 1 through 5
		philosophers[i] = &Philosopher{i + 1}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ { // only eat 3 times
			wg.Add(1)
			go philosophers[i].eat(host)
		}
	}

	host <- 1
	host <- 1

	wg.Wait()
}
