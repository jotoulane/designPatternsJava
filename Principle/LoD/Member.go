package LoD

type Member struct {
}

func NewMember() *Member {
	return &Member{}
}

/*
迪米特法则:对象之间的交互应该尽量简化，减少对象之间的耦合性
Boss只与TeamLeader直接交互，避免了和Member直接交互，减少了耦合性
*/
