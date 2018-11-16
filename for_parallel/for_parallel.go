// Sample program to show the basic concept of using a pointer
// to share data.
package main

import (
	"fmt"
	"time"
	"math/rand"
)

func printValue(value int){
	rt := rand.Int31n(1000)
	time.Sleep(time.Duration(rt) * time.Millisecond)
	fmt.Printf("valor = %v\r\n",value)
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

// main is the entry point for the application.
func main() {
	// Declare variable of type int with a value of 50.
	count := 50

	// Display the "value of" and "address of" count.
	println("Before:", count, &count)

	for i := 0; i<10; i++ {
		// Execute in parallel
		go printValue(i)
		// Execute in parallel - Pass the "address of" the variable count.
		go increment(&count)
		println("After: ", count, &count)
	}

	time.Sleep(1000 * time.Millisecond)
}
