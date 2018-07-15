package bll
type SendCommand struct {
Text string
To string
}

func (this SendCommand) From() IActor {
return nil
}