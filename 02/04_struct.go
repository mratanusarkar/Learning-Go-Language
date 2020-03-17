package main

import (
	"fmt"
)

type student struct {
	rollno uint16
	name   string
	marks  []float32
}

func (s student) avg() float32 {
	var sum float32 = 0
	for i := 0; i < len(s.marks); i++ {
		sum = sum + s.marks[i]
	}
	return sum / float32(len(s.marks))
}

func main() {
	stud1 = student {
		rollno: 1607022,
		name: "Atanu",
		marks: []float32{59.4, 90.0, 87.0, 89.0, 99.0, 100.0}
	}
}
