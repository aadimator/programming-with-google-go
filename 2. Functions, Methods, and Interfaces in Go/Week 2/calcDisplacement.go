// Write a program which first prompts the user to enter values for acceleration, 
// initial velocity, and initial displacement. Then the program should prompt the 
// user to enter a value for time and the program should compute the displacement 
// after the entered time.

package main

import (
	"fmt"
)

// GenDisplace computes displacement given the acceletation, initial velocity,
// initial displacement, and time.
func GenDisplace(a, v0, s0 float64) func (float64) float64 {
	return func(t float64) float64 {
		return  a*t*t/2 + v0*t + s0
	}
}

func main() {
	var a, v0, s0, t float64
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)

	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&v0)

	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&s0)

	fmt.Print("Enter time: ")
	fmt.Scan(&t)

	fn := GenDisplace(a, v0, s0)
	s := fn(t) // current displacement

	fmt.Printf("\nDisplacement would be %.2f meters after time  %.2f seconds.\n", s, t)
}