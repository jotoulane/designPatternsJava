package AbstractFactory

import "fmt"

type JavaVideo struct {
}

func (JavaVideo) produce() {
	fmt.Println("produce java")
}

type JavaArticle struct {
}

func (JavaArticle) produce() {
	fmt.Println("java 笔记")
}

type JavaCourseFactory struct {
}

func NewJavaCourseFactory() *JavaCourseFactory {
	return &JavaCourseFactory{}
}

func (JavaCourseFactory) getVideo() Video {
	return JavaVideo{}
}

func (JavaCourseFactory) getArticle() Article {
	return JavaArticle{}
}
