package bll

type SendCommand struct {
	Text string
	To   string
}

func (this SendCommand) From() string {
	return ""
}
