package bll

type Letter struct {
	Body      string
	Publisher Person
	To        string
}

func (this Letter) From() string {
	return this.Publisher.Name
}
