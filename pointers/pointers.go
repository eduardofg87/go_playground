// Sample program to show the basic concept of using a pointer
// to share data.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// Declare variable of type int with a value of 50.
	count := 50

	// Display the "value of" and "address of" count.
	println("Before:", count, &count)

	// Pass the "address of" the variable count.
	increment(&count)

	println("After: ", count, &count)
}

// increment declares count as a pointer variable whose value is
// always an address and points to values of type int.
func increment(inc *int) {
	// Increment the value that the "pointer points to". (de-referencing)
	*inc++
	println("Inc:   ", *inc, &inc, inc)

	// Do this to prevent inlining.
	var x int
	fmt.Sprintf("Prevent Inlining: %d", x)
}