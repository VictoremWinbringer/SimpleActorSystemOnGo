package bll

import "fmt"

type Person struct {
Name string
out  func(IMessage) error
}

func NewPerson(name string, out func(IMessage) error)  Person {
return Person{Name:name,out:out}
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