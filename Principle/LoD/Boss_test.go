package LoD

import "testing"

func TestBoss(t *testing.T) {
	boss := NewBoss()
	boss.commandCheckNumber(NewTeamLeader())
}
