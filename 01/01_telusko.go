package main // all go codes should be wrapped under a package

// fmt stands for format
import "fmt"

/* everything in Go happens inside functions */
func main2() {

	// fmt has all input/output functions
	fmt.Printf("using printf\n")
	fmt.Println("using println")

	// declaring a variable in go
	var num int = 3 // you can specify the type
	var a = 2       // or you can ignore it too!

	// make sure you use all declared variables and imported packages
	// else Go throws error
	fmt.Println(num + a + 5)

	// this will throw error, we need to specify type, or assign on declaration
	// var b

	var b int
	// do somthing
	b = 7 // you may assign later
	fmt.Println(b)

	// we can do this for multiple declaration
	var num1, num2 int
	fmt.Println(num1 + num2) // their default value is zero

	// we can assign in one line!
	num1, num2 = 2, 3
	fmt.Println(num1 + num2)

	result := 9 // shorthand for var result = 9
	fmt.Println(result)

	const pi = 3.14 // creating constant
	fmt.Println(add(pi, 0))

	sum, diff := calc(5, 4)
	fmt.Println(sum, diff)

	fmt.Println(operate(1, 1))

	// Loops in Go

	/*
	for{
		fmt.Println("im an infinite loop!")
	}
	*/

	fmt.Println("Loops in Go")
	i := 1
	for i <= 5 {
		fmt.Print(i)
		i++
	}

	fmt.Println()
	for j := 5; j >= 0; j-- {
		fmt.Print(j)
	}
}

// basic function structure
func add(x, y float32) float32 {
	return x + y
}

// we can return multiple variables
func calc(x, y int) (int, int) {
	return x + y, x - y
}

// matlab style maybe!
func operate(x, y int) (out1, out2 int) {
	out1 = 0
	out2 = 1
	return
}
