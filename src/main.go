package main

import (
	"fmt"
)

type IMessage interface {
	From() IActor
}

type Letter struct {
	Body string
	from Person
}

func (this Letter) From() IActor {
	return this.from
}

type SendCommand struct {
	Text string
}

func (this SendCommand) From() IActor {
	return nil
}

type IActor interface {
	In(message IMessage) error
	SetOut(func(IMessage) error) error
}

type Postman struct {
	out func(IMessage) error
}

func (this Postman) In(message IMessage) error {
	switch message.(type) {
	case Letter:
		return this.out(message)
	default:
		return fmt.Errorf("Unknown command %#v",message)
	}
}

func (this Postman) SetOut(handler func(IMessage) error) error{
	this.out = handler
	return nil
}

type Person struct {
	Name string
	out func(IMessage) error
}

func (this Person) In(message IMessage) error {
	switch message.(type) {
	case Letter:
		letter := message.(Letter)
		println(this.Name +" - recived message - "+ letter.Body)
		return nil
	case SendCommand:
		command := message.(SendCommand)
		return this.out(Letter{"From " +this.Name +" - "+command.Text, this})
	default:
		return fmt.Errorf("Unknown command %#v",message)
		}
}

func (this Person) SetOut(handler func(IMessage) error) error{
this.out = handler
return nil
}

func main() {
	var sender IActor = Person{"Foo", nil}
	var receiver IActor = Person{"Bar", nil}
	var postman IActor = Postman{ nil}
	var sendLetterCommand IMessage = SendCommand{"Hello World!"}

	sender.SetOut(func(message IMessage) error {
		return postman.In(message)
	})

	receiver.SetOut(func(message IMessage) error {
		println("Do nothing")
		return nil
	})

	postman.SetOut(func(message IMessage) error {
		return receiver.In(message)
	})

	sender.In(sendLetterCommand)
}
