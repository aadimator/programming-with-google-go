// Write a program which allows the user to create a set of animals and to get information 
// about those animals. Each animal has a name and can be either a cow, bird, or snake. 
// With each command, the user can either create a new animal of one of the three types, 
// or the user can request information about an animal that he/she has already created. 
// Each animal has a unique name, defined by the user. Note that the user can define animals 
// of a chosen type, but the types of animals are restricted to either cow, bird, or snake.

package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

// Animal defines the behaviors
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow is one specified animal
type Cow struct {
	name string
}

// Eat prints what a cow eats
func (a Cow) Eat() {
	fmt.Println(a.name + " eats: grass")
}

// Move prints how a cow moves
func (a Cow) Move() {
	fmt.Println(a.name + " moves: walk")
}

// Speak prints how a cow speaks
func (a Cow) Speak() {
	fmt.Println(a.name + " speaks: moo")
}

// Bird is one specified animal
type Bird struct {
	name string
}

// Eat prints how a bird eats
func (a Bird) Eat() {
	fmt.Println(a.name + " eats: worms")
}

// Move prints how a bird moves
func (a Bird) Move() {
	fmt.Println(a.name + " moves: fly")
}

// Speak prints how a bird speaks
func (a Bird) Speak() {
	fmt.Println(a.name + " speaks: peep")
}

// Snake is one specified animal
type Snake struct {
	name string
}

// Eat prints how a snake eats
func (a Snake) Eat() {
	fmt.Println(a.name + " eats: mice")
}

// Move prints how a snake moves
func (a Snake) Move() {
	fmt.Println(a.name + " moves: slither")
}

// Speak prints how a snake speaks
func (a Snake) Speak() {
	fmt.Println(a.name + " speaks: hsss")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	animals := make(map[string]Animal)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(input)
		s := strings.Split(strings.TrimSpace(input), " ")

		switch s[0] {
		case "newanimal":
			switch s[2] {
			case "cow":
				animals[s[1]] = Cow{s[1]}
			case "bird":
				animals[s[1]] = Bird{s[1]}
			case "snake":
				animals[s[1]] = Snake{s[1]}
			}

			fmt.Println("Created it!")
		case "query":
			a, ok := animals[s[1]]

			if ok {
				switch s[2] {
				case "eat":
					a.Eat()
				case "move":
					a.Move()
				case "speak":
					a.Speak()
				}
			} else {
				fmt.Println("Not found!")
			}
		}

	}
}