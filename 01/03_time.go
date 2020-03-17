package main

import (
	"fmt"
	"math"
	"reflect"
	"time"
)

func main() {

	var num float64 = 13
	var result = math.Sqrt(num)
	fmt.Printf("%.2f\n", result) //%g

	fmt.Println(reflect.TypeOf(time.Now().Weekday()))
}
