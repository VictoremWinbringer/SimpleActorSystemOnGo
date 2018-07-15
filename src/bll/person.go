package bll

import (
	"fmt"
)

type Person struct {
	Name      string
	receivers map[string]IActor
	attack    uint
	defence   int
	heals     uint
}

func NewPerson(name string) Person {
	return Person{Name: name, receivers: make(map[string]IActor), attack: 10, defence: 5, heals: 100}
}

func (this Person) SetReceiver(receiver IActor, id string) error {
 this.receivers[id] = receiver
 return nil
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
	case AttackCommand:
		command := message.(AttackCommand)
		return this.outAttack(command)
	case TakeDamageCommand:
		command := message.(TakeDamageCommand)
		fmt.Printf("%s damaged by %s  value %d \n",this.Name, command.From(), command.Value)
		this.takeDamage(command)
		fmt.Printf("Current heals %d\n", this.heals)
		return nil
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

func (this Person) outAttack(message AttackCommand) error {
	target, ok := this.receivers[message.Target]
	if !ok {
		return fmt.Errorf("can't find target")
	}
	return target.In(TakeDamageCommand{this.Name, this.attack})
}

func (this *Person) takeDamage(message TakeDamageCommand) {
	damage := int(message.Value) - this.defence
	if damage < 1 {
		return
	}
	if damage >= int(this.heals) {
		this.heals = 0
		return
		}
	this.heals -= uint(damage)
}
