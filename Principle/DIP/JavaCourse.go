package DIP

import "fmt"

type JavaCourse struct {
}

func NewJavaCourse() *JavaCourse {
	return &JavaCourse{}
}

func (JavaCourse) studyCourse() {
	fmt.Println("study java")
}
