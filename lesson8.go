/*
Let us assume the following formula for displacement s as a function of time t,
acceleration a, initial velocity vo, and initial displacement so.
s = ½ a t2 + vot + so

Write a program which first prompts the user to enter values for acceleration,
initial velocity, and initial displacement. Then the program should prompt the user
to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments,
acceleration a, initial velocity vo, and initial displacement so.

GenDisplaceFn() should return a function which computes displacement as a function of time, assuming
the given values acceleration, initial velocity, and initial displacement.
The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and
return one float64 argument which is the displacement travelled after time t.

For example, let's say that I want to assume the following values for acceleration,
initial velocity, and initial displacement: a = 10, vo = 2, so = 1.
I can use the following statement to call GenDisplaceFn() to generate a function
fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)
Then I can use the following statement to print the displacement after 3 seconds.
fmt.Println(fn(3))
And I can use the following statement to print the displacement after 5 seconds.
fmt.Println(fn(5))

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// Prompt the user to enter values for acceleration, initial velocity, and initial displacement.
	var a, vo, so float64
	fmt.Print("Acceleration: ")
	fmt.Scanln(&a)
	fmt.Print("Initial velocity: ")
	fmt.Scanln(&vo)
	fmt.Print("Initial displacement: ")
	fmt.Scanln(&so)

	// Generate the displacement function using GenDisplaceFn.
	fn := GenDisplaceFn(a, vo, so)

	// Prompt the user to enter a value for time.
	var t float64
	fmt.Print("Time: ")
	fmt.Scanln(&t)

	// Compute the displacement after the entered time and print it.
	displacement := fn(t)
	fmt.Println("Displacement after", t, "seconds:", displacement)
}

// GenDisplaceFn generates a displacement function based on the given acceleration,
// initial velocity, and initial displacement.
func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + vo*t + so
	}
}
