package FactoryMethod

type JavaVideoFactory struct {
}

func NewJavaFactory() *JavaVideoFactory {
	return &JavaVideoFactory{}
}

func (j JavaVideoFactory) getVideo() Video {
	return JavaVideo{}
}
