package FactoryMethod

type VideoFactory interface {
	getVideo() Video
}

/*
通过定义一个创建对象的接口，但是让子类决定实例化哪个类
*/
