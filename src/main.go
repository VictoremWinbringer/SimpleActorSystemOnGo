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
	To   string
}

func (this Letter) From() IActor {
	return &this.from
}

type SendCommand struct {
	Text string
	To string
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
		println("Postman process letter from: " + message.From().(*Person).Name)
		return this.out(message)
	default:
		return fmt.Errorf("Unknown command %#v", message)
	}
}

func (this *Postman) SetOut(handler func(IMessage) error) error {
	this.out = handler
	return nil
}

type Person struct {
	Name string
	out  func(IMessage) error
}

func (this Person) In(message IMessage) error {
	switch message.(type) {
	case Letter:
		letter := message.(Letter)
		println(this.Name + " - recived message - " + letter.Body)
		return nil
	case SendCommand:
		command := message.(SendCommand)
		return this.out(Letter{"From " + this.Name + " - " + command.Text, this, command.To})
	default:
		return fmt.Errorf("Unknown command %#v", message)
	}
}

func (this *Person) SetOut(handler func(IMessage) error) error {
	this.out = handler
	return nil
}

func main() {
	var sender IActor = &Person{"Foo", nil}
	var bar IActor = &Person{"Bar", nil}
	var baz IActor = &Person{"Baz", nil}
	var postman IActor = &Postman{nil}

	sender.SetOut(func(message IMessage) error {
		return postman.In(message)
	})

	bar.SetOut(func(message IMessage) error {
		println("Do nothing")
		return nil
	})

	baz.SetOut(func(message IMessage) error {
		println("Do nothing")
		return nil
	})

	postman.SetOut(func(message IMessage) error {
		switch message.(type) {
		case Letter:
			to := message.(Letter).To
			switch to {
			case "Bar":
				return bar.In(message)
			case "Baz":
				return baz.In(message)
			default:
				return fmt.Errorf("Dont now name " + to)
			}
		default:
			return fmt.Errorf("Uknown message %#v",message)
		}
	})
	
	sender.In(SendCommand{"Hello World!", "Bar"})
	sender.In(SendCommand{"It's work! :)", "Baz"})
	fmt.Scanln()
}
