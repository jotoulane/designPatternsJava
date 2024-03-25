package OCP

type JavaCourse struct {
	id    int
	name  string
	price float32
}

func NewJavaCourse(id int, name string, price float32) *JavaCourse {
	return &JavaCourse{id: id, name: name, price: price}
}

func (j *JavaCourse) GetID() int {
	return j.id
}

func (j *JavaCourse) GetName() string {
	return j.name
}

func (j *JavaCourse) GetPrice() float32 {
	return j.price
}
