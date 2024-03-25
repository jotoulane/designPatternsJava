package DIP

import "fmt"

type PythonCourse struct {
}

func NewPythonCourse() *PythonCourse {
	return &PythonCourse{}
}

func (PythonCourse) studyCourse() {
	fmt.Println("study python")
}
