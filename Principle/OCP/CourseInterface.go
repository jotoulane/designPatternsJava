package OCP

type CourseInterface interface {
	GetID() int
	GetName() string
	GetPrice() float32
}

//开闭原则：对扩展开放，对修改关闭
//通过继承，保持了原有类的封闭添加了新的功能
