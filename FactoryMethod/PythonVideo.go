package FactoryMethod

import "fmt"

type PythonVideo struct {
}

func (PythonVideo) produce() {
	fmt.Println("produce python")
}
