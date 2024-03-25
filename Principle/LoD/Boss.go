package LoD

type Boss struct {
}

func NewBoss() *Boss {
	return &Boss{}
}

func (Boss) commandCheckNumber(leader *TeamLeader) {
	leader.checkNumberOfMember()
}
