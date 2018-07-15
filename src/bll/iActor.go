package bll

type IActor interface {
	In(message IMessage) error
}
