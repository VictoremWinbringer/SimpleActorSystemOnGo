package bll

import "fmt"

type Postman struct {
out func(IMessage) error
}

func NewPostman(out func(IMessage) error)  Postman {
	return Postman{out:out}
}

func (this Postman) In(message IMessage) error {
switch message.(type) {
case Letter:
println("Postman process letter from: " + message.From().(*Person).Name)
return this.out(message)
default:
return fmt.Errorf("Unknown command %#v", message)
}
}