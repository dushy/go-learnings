/*
Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters 'i', 'a', and 'n'.
The program should print "Found!" if the entered string starts with the character 'i',
ends with the character 'n', and contains the character 'a'.

The program should print "Not Found!" otherwise.

The program should not be case-sensitive, so it does not matter if the characters are
upper-case or lower-case.
Examples: The program should print "Found!" for the following example entered strings,
"ian", "Ian", "iuiygaygn", "I d skd a efju N".

The program should print "Not Found!" for the following strings, "ihhhhhn", "ina", "xian".

Submit your source code for the program, "findian.go".
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter a string:")
	var input string
	fmt.Scan(&input)

	if strings.HasPrefix(strings.ToLower(input), "i") && strings.HasSuffix(strings.ToLower(input), "n") && strings.Contains(strings.ToLower(input), "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
