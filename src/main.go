package main

import (
	"fmt"
	. "./bll"
)

func main() {

	var bar IActor = NewPerson("Bar", func(message IMessage) error {
		println("Do nothing")
		return nil
	})
	var baz IActor = NewPerson("Baz", func(message IMessage) error {
		println("Do nothing")
		return nil
	})

	var postman IActor = NewPostman(func(message IMessage) error {
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
			return fmt.Errorf("Uknown message %#v", message)
		}
	})

	var sender IActor = NewPerson("Foo", func(message IMessage) error {
		return postman.In(message)
	})

	sender.In(SendCommand{"Hello World!", "Bar"})
	sender.In(SendCommand{"It's work! :)", "Baz"})
	fmt.Scanln()
}
