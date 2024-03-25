package LoD

import (
	"container/list"
	"fmt"
)

type TeamLeader struct {
}

func NewTeamLeader() *TeamLeader {
	return &TeamLeader{}
}

func (TeamLeader) checkNumberOfMember() {
	l := list.New()
	l.Init()
	for i := 0; i < 20; i++ {
		l.PushBack(NewMember())
	}
	fmt.Println(l.Len())
}
