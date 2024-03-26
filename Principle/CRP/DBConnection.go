package CRP

type DBConnection interface {
	GetConnection() string
}

/*
合成复用原则
使用组合或者聚合来实现代码复用，而不是继承
*/
