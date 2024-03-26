package FactoryMethod

import "fmt"

type JavaVideo struct {
}

func (j JavaVideo) produce() {
	fmt.Println("produce java")
}
