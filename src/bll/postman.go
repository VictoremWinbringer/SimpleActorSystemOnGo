package bll

import "fmt"

type Postman struct {
	receivers map[string]IActor
}

func NewPostman() Postman {
	return Postman{receivers: make(map[string]IActor)}
}

func (this Postman) SetReceiver(receiver IActor, id string) error {
	this.receivers[id] = receiver
	return nil
}

func (this Postman) In(message IMessage) error {
	switch message.(type) {
	case Letter:
		println("Postman process letter from: " + message.From())
		return this.out(message)
	default:
		return fmt.Errorf("Unknown command %#v", message)
	}
}

func (this Postman) out(message IMessage) error {
	switch message.(type) {
	case Letter:
		to := message.(Letter).To
		receiver, ok := this.receivers[to]
		if !ok {
			return fmt.Errorf("Dont now name " + to)
		}
		return receiver.In(message)
	default:
		return fmt.Errorf("Uknown message %#v", message)
	}
}
