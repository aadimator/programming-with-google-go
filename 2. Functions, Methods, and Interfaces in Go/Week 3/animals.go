// Write a program which allows the user to get information about a predefined set of animals. 
// Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak. 
// The user can issue a request to find out one of three things about an animal: 
// 1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks. 

package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

type Animal struct {
	food, locomotion, noise string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(input)
		s := strings.Split(strings.TrimSpace(input), " ")

		a := new(Animal)
		switch s[0] {
		case "cow":
			a.food = "grass"
			a.locomotion = "walk"
			a.noise = "moo"
		case "bird":
			a.food = "worms"
			a.locomotion = "fly"
			a.noise = "poop"
		case "snake":
			a.food = "mice"
			a.locomotion = "slither"
			a.noise = "hsss"
		}

		switch s[1] {
		case "eat":
			a.Eat()
		case "speak":
			a.Speak()
		case "move":
			a.Move()
		}
	}
}