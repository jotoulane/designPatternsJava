package AbstractFactory

type Article interface {
	produce()
}

type Video interface {
	produce()
}

type CourseFactory interface {
	getVideo() Video
	getArticle() Article
}
