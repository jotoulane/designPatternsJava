package DIP

type CourseInterface interface {
	studyCourse()
}

//依赖倒置原则：高层模块不应该依赖于低层模块，而都应该依赖抽象
//Student类不依赖于具体的课程实现，而是依赖于CourseInterface接口
