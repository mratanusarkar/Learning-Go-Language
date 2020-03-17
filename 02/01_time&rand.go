package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*TimeNRand is a function containing
 * basic learning stuffs with time packages
 * basic way to generate random numbers
 */
func TimeNRand() {
	// Seeding with the same value results in the same random sequence each run.
	// For different numbers, seed with a different value, such as
	// time.Now().UnixNano(), which yields a constantly-changing number.

	rand.Seed(time.Now().UnixNano())
	fmt.Println("generating 3 random numbers:")
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println()

	iterations := 9999999999

	fmt.Println("Time counter started!")
	var t0 = time.Now()
	var t0n int64 = time.Now().UnixNano()

	fmt.Println("Loop running with number of iterations =", iterations)
	for i := 0; i < iterations; i++ {

	}

	var t1 = time.Now()
	var t1n int64 = time.Now().UnixNano()

	fmt.Println("Time counter stopped!")
	fmt.Println("Time taken:")

	fmt.Println(t1.Sub(t0))
	fmt.Println(t1n - t0n)
}
