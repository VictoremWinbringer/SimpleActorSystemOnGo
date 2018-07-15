package bll

import "fmt"

type Person struct {
	Name      string
	receivers map[string]IActor
}

func NewPerson(name string, receivers map[string]IActor) Person {
	return Person{Name: name, receivers: receivers}
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

func (this Person) out(message Letter) error {
	postman, ok := this.receivers["postman"]
	if !ok {
		return fmt.Errorf("can't find postman")
	}
	return postman.In(message)
}
