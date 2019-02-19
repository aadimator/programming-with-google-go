// Write a program which prompts the user to first enter a name, and then enter an address. 
// Your program should create a map and add the name and address to the map using the keys 
// “name” and “address”, respectively. Your program should use Marshal() to create a JSON object 
// from the map, and then your program should print the JSON object.

package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	m := make(map[string]string)

	var name, address string

	fmt.Print("Enter name: ")
	fmt.Scan(&name)
	fmt.Print("Enter address: ")
	fmt.Scan(&address)

	m["name"] = name
	m["address"] = address

	json, _ := json.Marshal(m)
	
	fmt.Println()
	fmt.Println("JSON")
	fmt.Println(string(json))
}