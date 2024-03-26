package AbstractFactory

import "testing"

func TestAbstractFactory(t *testing.T) {

	var courseFactory CourseFactory = NewJavaCourseFactory()
	courseFactory.getVideo().produce()
	courseFactory.getArticle().produce()

}
