/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys
"name" and "address", respectively.

Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.

Submit your source code for the program, "makejson.go".
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string

	fmt.Print("Enter a name: ")
	fmt.Scan(&name)
	fmt.Print("Enter an address: ")
	fmt.Scan(&address)

	m := map[string]string{
		"name":    name,
		"address": address,
	}

	json, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(json))
}
