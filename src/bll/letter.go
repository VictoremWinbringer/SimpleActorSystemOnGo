package bll

type Letter struct {
	Body string
	Publisher Person
	To   string
}

func (this Letter) From() IActor {
	return &this.Publisher
}