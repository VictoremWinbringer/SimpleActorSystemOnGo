package bll

type AttackCommand struct {
	Target string
}

func (this AttackCommand) From() string {
	return ""
}
