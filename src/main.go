package main

import (
	"fmt"
	. "./bll"
)

func main() {

	var bar IActor = NewPerson("Bar", make(map[string]IActor,0))
	var baz IActor = NewPerson("Baz",make(map[string]IActor,0))
	 postmanReceivers := make(map[string]IActor,0)
	 postmanReceivers["Bar"] = bar
	 postmanReceivers["Baz"] = baz
	var postman IActor = NewPostman(postmanReceivers)
	senderReceivers := make(map[string]IActor,0)
	senderReceivers["postman"] = postman
	var sender IActor = NewPerson("Foo",senderReceivers)
	e:=sender.In(SendCommand{"Hello World!", "Bar"})
   printError(e)
 	e = sender.In(SendCommand{"It's work! :)", "Baz"})
	printError(e)
	fmt.Scanln()
}

func printError(e error){
	if e != nil {
		println(e)
	}
}
