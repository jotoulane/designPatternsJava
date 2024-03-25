package ISP

import "fmt"

type Dog struct {
}

func (Dog) eat() {
	fmt.Println("bird eat")
}

func (Dog) swim() {
	fmt.Println("bird eat")
}
