package OCP

type JavaDiscountCourse struct {
	JavaCourse
}

func NewJavaDiscountCourse(ID int, name string, price float32) *JavaDiscountCourse {
	return &JavaDiscountCourse{JavaCourse: *NewJavaCourse(ID, name, price)}
}

func (j *JavaDiscountCourse) GetOriginPrice() float32 {
	return j.GetPrice()
}

func (j *JavaDiscountCourse) GetDiscountPrice() float32 {
	return j.GetPrice() * 0.8
}
