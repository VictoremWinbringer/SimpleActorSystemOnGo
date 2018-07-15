package bll

type IActor interface {
	In(message IMessage) error
	SetReceiver(receiver IActor, id string) error
}
