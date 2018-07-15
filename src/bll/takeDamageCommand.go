package bll

type TakeDamageCommand struct {
	Sender string
	Value  uint
}

func (this TakeDamageCommand) From() string {
	return this.Sender
}
