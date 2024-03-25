package DIP

type Student struct {
	course CourseInterface
}

func NewStudent(course CourseInterface) *Student {
	return &Student{course: course}
}

func (s Student) studyCourse() {
	s.course.studyCourse()
}
