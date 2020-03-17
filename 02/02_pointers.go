package main

import "fmt"

// LearnPointers is a function describing syntax and usage of pointers in Go
func LearnPointers() {

	// container (variable/memory area) of type int
	var num int = 10

	// a pointer can only store address! (container storing integer address)
	var ptr *int = &num

	fmt.Println("--------------------------------")
	fmt.Println("ptr[add of num] -> num[ int 10 ]")
	fmt.Println("--------------------------------")
	fmt.Println()

	fmt.Println("value inside num:", num) // value inside num
	fmt.Println("address of num:", &num)  // address of num
	//fmt.Println("value pointed by num", *num)	// it's not a pointer! it can't point to somthing!
	fmt.Println()

	fmt.Println("value inside ptr:", ptr)      // value inside ptr
	fmt.Println("address of ptr:", &ptr)       // address of ptr
	fmt.Println("value pointed by ptr:", *ptr) // value of the stuff it's pointing to
	fmt.Println("--------------------------------")

}
