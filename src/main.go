package main

import (
	"fmt"
	. "./bll"
)

func main() {

	var bar IActor = NewPerson("Bar")
	var baz IActor = NewPerson("Baz")
	var sender IActor = NewPerson("Foo")
	var postman IActor = NewPostman()
	sender.SetReceiver(bar, "Bar")
	sender.SetReceiver(baz, "Baz")
	sender.SetReceiver(postman, "postman")
	postman.SetReceiver(bar, "Bar")
	postman.SetReceiver(baz, "Baz")
	e := sender.In(SendCommand{"Hello World!", "Bar"})
	printError(e)
	e = sender.In(SendCommand{"It's work! :)", "Baz"})
	printError(e)
	e = sender.In(AttackCommand{"Bar"})
	printError(e)
	fmt.Scanln()
}

func printError(e error) {
	if e != nil {
		println(e.Error())
	}
}
