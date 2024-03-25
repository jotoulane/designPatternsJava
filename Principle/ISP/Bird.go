package ISP

import "fmt"

type Bird struct {
}

func (Bird) fly() {
	fmt.Println("bird fly")
}

func (Bird) eat() {
	fmt.Println("bird eat")
}
