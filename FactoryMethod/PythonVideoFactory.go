package FactoryMethod

type PythonVideoFactory struct {
}

func NewPythonFactory() *PythonVideoFactory {
	return &PythonVideoFactory{}
}

func (PythonVideoFactory) getVideo() Video {
	return PythonVideo{}
}
