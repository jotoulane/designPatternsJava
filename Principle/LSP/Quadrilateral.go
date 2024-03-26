package LSP

type Quadrilateral interface {
	Width() int
	Length() int
}

/*
里氏替换原则
所有引用父类的地方都能使用子类对象
*/
