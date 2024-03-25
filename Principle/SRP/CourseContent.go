package SRP

type CourseContent interface {
	getCourseName() string
	getCourseVideo() []byte
}

//单一职责原则:一个类或模块应该只有一个责任
//将不同的功能分离到不同的接口和结构体中，从而遵循了单一职责原则
