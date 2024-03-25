package SRP

import "fmt"

type Course struct {
}

func (Course) getCourseName() string {
	return "course name"
}

func (Course) getCourseVideo() []byte {
	return nil
}

func (Course) studyCourse() {
	fmt.Println("study")
}
func (Course) refundCourse() {
	fmt.Println("refund")
}
