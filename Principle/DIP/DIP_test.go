package DIP

import "testing"

func TestDIP(t *testing.T) {
	student := NewStudent(NewPythonCourse())
	student.studyCourse()
}
