package internal

type Subscriber struct {
	Name    string
	Surname string
	Tel     string
}

func NewSubscriber(name, surname, tel string) *Subscriber {
	return &Subscriber{Name: name, Surname: surname, Tel: tel}
}
